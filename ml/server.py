from flask import Flask
from influxdb_client import InfluxDBClient
from apscheduler.schedulers.background import BackgroundScheduler
from model import BILSTM
from collections import defaultdict
import numpy as np
from requests import get, post
import time
import math

app = Flask(__name__)

my_url = 'http://152.7.179.7:8086/'
my_org = 'f87217d97d859651'
my_uname = 'admin'
my_passwd = 'password'

client = InfluxDBClient(url=my_url, org=my_org, username=my_uname, password=my_passwd)
bucket = 'TrainingData'
tag = '_measurement'
tag_val = 'metrics'

currReplicas = 1

model = BILSTM()

def get_new_data():
    cpu_mem_tables = client.query_api().query('from(bucket:"TrainingData") |> range(start: -3m) \
                                            |> filter(fn: (r) => r._measurement == "metrics" and r.type == "agrigate" and (r._field == "mem" or r._field == "cpu")) \
                                            |> truncateTimeColumn(unit: 1s) \
                                            |> sort(columns: ["_time"]) \
                                            |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")')
    
    rps_tables = client.query_api().query('from(bucket:"TrainingData") |> range(start: -3m) \
                                        |> filter(fn: (r) => r._measurement == "metrics" and r.type == "agrigate" and (r._field == "RPS")) \
                                        |> truncateTimeColumn(unit: 1s) \
                                        |> sort(columns: ["_time"]) \
                                        |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")')

    result = defaultdict(dict)

    for table in cpu_mem_tables:
        for row in table:
            result[row["_time"]]['cpu'] = row['cpu']
            result[row["_time"]]['mem'] = row['mem']
    for table in rps_tables:
        for row in table:
            result[row["_time"]]['RPS'] = row['RPS']

    return np.array([[result[entry]['cpu'], result[entry]['RPS'], result[entry]['mem']] for entry in sorted(result)])

def get_train_data():
    cpu_mem_tables = client.query_api().query('from(bucket:"TrainingData") |> range(start: -5h) \
                                            |> filter(fn: (r) => r._measurement == "metrics" and r.type == "agrigate" and (r._field == "mem" or r._field == "cpu")) \
                                            |> truncateTimeColumn(unit: 1s) \
                                            |> sort(columns: ["_time"]) \
                                            |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")')
    
    rps_tables = client.query_api().query('from(bucket:"TrainingData") |> range(start: -5h) \
                                        |> filter(fn: (r) => r._measurement == "metrics" and r.type == "agrigate" and (r._field == "RPS")) \
                                        |> truncateTimeColumn(unit: 1s) \
                                        |> sort(columns: ["_time"]) \
                                        |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")')

    result = defaultdict(dict)

    for table in cpu_mem_tables:
        for row in table:
            result[row["_time"]]['cpu'] = row['cpu']
            result[row["_time"]]['mem'] = row['mem']
    for table in rps_tables:
        for row in table:
            result[row["_time"]]['RPS'] = row['RPS']

    return np.array([[result[entry]['cpu'], result[entry]['RPS'], result[entry]['mem']] for entry in sorted(result)])


def scale_containers(pdata, cdata):
    global currReplicas
    #add formula to scale containers
    print("Scaling containers")
    print(pdata, cdata)
    predectedCPU = pdata[0][0]
    currentCPU = cdata
    response = get('http://152.7.179.7:8000/admin/getcontainerscount')
    currReplicas = int(response.text)
    desiredReplicas = max(min(5, math.ceil(currReplicas * (predectedCPU/50))), 1)
    s = "" + str(desiredReplicas) + " " + str(predectedCPU) + "\n"
    file = open('Prediction.txt', 'a+')
    file.write(s)
    x = post('http://152.7.179.7:8000/admin/updatecontainers', json = {
    "app_name":"rubis",
    "container_num": desiredReplicas
    })

def process_data(data):
    try:
        reshaped_data = data.reshape((data.shape[0]//4, 4, 3))
        avg_data = np.mean(reshaped_data, axis=1)
        return avg_data
    except:
        print("Invalid data shape")
        return None
    

def monitor():
    global currReplicas
    print('Fetching Influx DB data...')
    data = get_new_data()
    data = process_data(data)
    if data is None:
        return
    cpu = model.predict(data, currReplicas)
    scale_containers(cpu, data[-1][0])
    print(cpu)

def train():
    print('Fetching Influx DB data for training...')
    data = get_train_data()
    data = process_data(data)
    if data is None:
        return
    model.train(data)

scheduler = BackgroundScheduler()
scheduler.add_job(monitor, 'interval', minutes=1)
scheduler.start()

scheduler_t = BackgroundScheduler()
scheduler_t.add_job(train, 'interval', minutes=300)

if __name__ == '__main__':
    app.run()
