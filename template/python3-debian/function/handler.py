import requests

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    outputStr="f1,"+"Start"
    requests.post("http://172.22.85.31:30036",outputStr)
    return "Sent request to the coordinator from f1"