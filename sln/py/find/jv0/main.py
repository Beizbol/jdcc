from bottle import request, route, run, static_file


@route("/")
def home():
    return static_file("find.html", root="../../../../www/")


@route("/find", method="POST")
def find():
    search = request.forms.get("search")
    # TODO: Implement search
    return f"Searching for {search}..."


if __name__ == "__main__":
    run(host="localhost", port=4269)
