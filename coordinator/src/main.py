from fastapi import FastAPI

from repos.db import get_db

app = FastAPI()


@app.get("/")
async def root():
    a = get_db()
    return {"message": str(a)}