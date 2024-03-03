from bottle import route, run


@route("/hey")
def index(name):
    return "Hey from Python!"


run(host="localhost", port=4269)
