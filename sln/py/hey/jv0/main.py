from blacksheep import Application, get
import uvicorn as server

app = Application()


@get("/")
async def home():
    return "Hello, World!"


@get("/hey")
async def hey():
    return "Hey from Python!"


if __name__ == "__main__":
    server.run("main:app", port=4269, log_level="info")
