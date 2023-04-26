from flask import Flask
from apscheduler.schedulers.background import BackgroundScheduler
from model import BILSTM
from collections import defaultdict
import time

app = Flask(__name__)

my_url = 'http://152.7.179.7:8086/'
my_org = 'f87217d97d859651'
my_uname = 'admin'
my_passwd = 'password'

client = InfluxDBClient(url=my_url, org=my_org, username=my_uname, password=my_passwd)
bucket = 'TrainingData'
tag = '_measurement'
tag_val = 'metrics'

def get_new_data():
    cpu_mem_tables =    client.query_api().query('from(bucket:"TrainingData") |> range(start: -10m) \
                                                |> filter(fn: (r) => r._measurement == "metrics" and r.type == "agrigate" and (r._field == "mem" or r._field == "cpu")) \
                                                |> truncateTimeColumn(unit: 1s) \
                                                |> sort(columns: ["_time"]) \
                                                |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")')
    
    rps_tables =        client.query_api().query('from(bucket:"TrainingData") |> range(start: -10m) \
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

    return numpy.array([(entry, (result[entry]['cpu'], result[entry]['RPS'], result[entry]['mem'])) for entry in sorted(result)])

def convert_tables_to_dict(tables):
    data = defaultdict(dict)
    for table in tables:
        for row in table:
            data[row['_time']][row['_field']] = row['_value']
    return data


def scale_containers(data):
    pass


def monitor():
    print('Fetching Influx DB data...')
    data = get_data()


scheduler = BackgroundScheduler()
scheduler.add_job(monitor, 'interval', minutes=5)
scheduler.start()

if __name__ == '__main__':
    app.run()
