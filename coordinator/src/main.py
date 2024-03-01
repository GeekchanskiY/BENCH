import aiohttp

from fastapi import FastAPI

from routers.users import router as user_router

from models.db import get_db, get_redis

app = FastAPI()

app.include_router(user_router)

@app.get("/")
async def root():
    a = get_db()
    return {"message": str(a),
            "ping": "pong3"}

@app.get('/test')
async def test():
    resp_text = ''
    async with aiohttp.ClientSession() as session:
        async with session.get('http://ressearch_backend:3000/ping') as resp:
            # print(resp.status)
            resp_text = await resp.json()
    
    redis = await get_redis()
    await redis.set('test', 'hello, redis')
    redis_value = await redis.get('test')

    return {
        'ruby response': resp_text,
        'redis response': redis_value
    }