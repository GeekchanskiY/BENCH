import aiohttp

from fastapi import FastAPI

from routers.users import router as user_router

from fastapi.middleware.cors import CORSMiddleware

from models.db import get_db, get_redis

import pika

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
    credentials = pika.PlainCredentials('rmuser', 'rmpassword')
    parameters = pika.ConnectionParameters('rabbitmq',
                                    5672,
                                    '/',
                                    credentials)
    connection = pika.BlockingConnection(parameters=parameters)
    channel = connection.channel()
    channel.queue_declare(queue='hello')

    channel.basic_publish(exchange='',
                    routing_key='hello',
                    body='Hello W0rld!')
    connection.close()
    ressearch_response = 'Failed'
    redis_response = 'Failed'
    finance_response = 'Failed'
    try:
        async with aiohttp.ClientSession() as session:
            async with session.get('http://ressearch_backend:3000/ping') as resp:
                val = await resp.json()
                if val.get('msg', None) == 'pong':
                    ressearch_response = 'Success'
    except Exception as e:
        ressearch_response = str(e)
    
    try:
        redis = await get_redis()
        await redis.set('test', 'test')
        redis_value = await redis.get('test')
        if redis_value.decode('utf-8') == 'test':
            redis_response = 'Success'
    except Exception as e:
        redis_response = e.text

    try:
        async with aiohttp.ClientSession() as session:
            async with session.get('http://finance_backend:3001/v1/ping') as resp:
                val = await resp.json()
                if val.get('message', None) == 'pong':
                    finance_response = 'Success'
    except Exception as e:
        finance_response = str(e)


    return {
        'Ressearch': ressearch_response,
        'Redis': redis_response,
        'Finance': finance_response
    }