import aiohttp

from fastapi import FastAPI

from models.db import get_db, get_redis

app = FastAPI()


@app.get("/")
async def root():
    a = get_db()
    return {"message": str(a)}

@app.get('/test')
async def test():
    resp_text = ''
    async with aiohttp.ClientSession() as session:
        async with session.get('http://ressearch_backend:3000/') as resp:
            # print(resp.status)
            resp_text = await resp.text()
    
    redis = await get_redis()
    await redis.set('test', 'hello, redis')
    redis_value = await redis.get('test')

    return {
        'ruby response': resp_text,
        'redis response': redis_value
    }