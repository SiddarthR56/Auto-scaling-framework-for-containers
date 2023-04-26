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

def get_data():
    tables = client.query_api().query(f'from(bucket:"{bucket}") |> range(start: {start}) \
                                                                |> filter(fn: (r) => r.{tag} == "{tag_val}" and r.id == "{container}") \
                                                                |> truncateTimeColumn(unit: 1s) \
                                                                |> sort(columns: ["_time"])')

    data = convert_tables_to_dict(tables)
    return [(entry, (data[entry]['cpuUsage'], data[entry]['memUsage'])) for entry in sorted(data)]


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
