from blacksheep import Application as App, ContentDispositionType
from blacksheep import Request, get, post, file
import uvicorn as server

app = App()


@get("/")
async def home():
    return file(
        "../../../../www/find.html",
        content_type="text/html",
        content_disposition=ContentDispositionType.INLINE,
    )


@get("/hey")
async def hey():
    return "Hey from Python!"


@post("/find")
async def find(request: Request):
    form_dict = await request.form()
    search = form_dict["search"]
    return f"Searching for: {search}"


if __name__ == "__main__":
    server.run("main:app", port=4269, log_level="info")
