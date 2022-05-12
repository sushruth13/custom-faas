import requests

def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """
    testvar=""
    #pass the variables from the function to be sent to the next function in wf. Please remeber to pass the current fn name
    requests.post("http://coordinator-svc.default.svc.cluster.local:80",testvar)

    return req
