import aiohttp

from fastapi import FastAPI

from routers.users import router as user_router

from fastapi.middleware.cors import CORSMiddleware

from models.db import get_db, get_redis

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['http://0.0.0.0:3002'],
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*']
)

app.include_router(user_router)

@app.get("/")
async def root():
    a = get_db()
    return {"message": str(a),
            "ping": "pong3"}

@app.get('/healthcheck')
async def healthcheck():
    ressearch_response = 'Failed'
    redis_response = 'Failed'
    finance_response = 'Failed'
    async with aiohttp.ClientSession() as session:
        async with session.get('http://ressearch_backend:3000/ping') as resp:
            val = await resp.json()
            if val.get('msg', None) == 'pong':
                ressearch_response = 'Success'
    
    try:
        redis = await get_redis()
        await redis.set('test', 'test')
        redis_value = await redis.get('test')
        if redis_value.decode('utf-8') == 'test':
            redis_response = 'Success'
    except Exception as e:
        redis_response = e.text

    async with aiohttp.ClientSession() as session:
        async with session.get('http://finance_backend:3001/v1/ping') as resp:
            val = await resp.json()
            if val.get('message', None) == 'pong':
                finance_response = 'Success'
  


    return {
        'Ressearch': ressearch_response,
        'Redis': redis_response,
        'Finance': finance_response
    }