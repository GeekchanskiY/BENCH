import aiohttp

from fastapi import FastAPI

from repos.db import get_db

app = FastAPI()


@app.get("/")
async def root():
    a = get_db()
    return {"message": str(a)}

@app.get('/test')
async def test():
    async with aiohttp.ClientSession() as session:
        async with session.get('http://ressearch_backend:3000/') as resp:
            # print(resp.status)
            return{resp.status: await resp.text()}