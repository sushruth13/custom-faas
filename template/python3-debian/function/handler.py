import requests

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    emptyvar=""
    requests.post("http://coordinator-svc.default.svc.cluster.local:80",emptyvar)
    return req
