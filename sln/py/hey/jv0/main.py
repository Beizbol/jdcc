from bottle import route, run


@route("/")
@route("/hey")
def hey():
    return "Hey from Python!"


if __name__ == "__main__":
    run(host="localhost", port=4269)
