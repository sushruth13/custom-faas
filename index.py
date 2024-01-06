import requests
import uuid
from time import sleep
import random


def main():
    sleep_duration = random.uniform(5, 10)
    sleep(sleep_duration)
    #sending postrequest to recorder
    recorder_url="http://172.22.85.201:31082/record"

    #sends data to the recorder
    data=requests.post(recorder_url,combined_data)

    #sending data to Orachestrator


    orachestrator_url="http://172.22.85.201:30578/orachestra"
    data=requests.post(orachestrator_url,combined_data)

