import requests
from time import sleep
import random

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    sleep(random.randint(1,10))
    
    outputStr="wf1-f1,"+"Start"
    requests.post("http://172.22.85.31:30036",outputStr)
    #data=requests.post("http://172.22.85.31:31112/function/f2","1")
    return "Sent the request to coordinator from wf1-f1"
