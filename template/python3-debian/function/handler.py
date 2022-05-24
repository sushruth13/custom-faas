import requests

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    requests.post("http://coordinator.default.svc.cluster.local:8080")
    return req
