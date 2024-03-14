import aiohttp

from fastapi import FastAPI, WebSocket, WebSocketDisconnect

from routers.users import router as user_router
from routers.services import router as service_router
from routers.sockets import router as socket_router

# from fastapi.middleware.cors import CORSMiddleware
from starlette.middleware.cors import CORSMiddleware

from repositories.models.db import get_db, get_redis
import aiokafka
import pika


app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*', 'http://0.0.0.0:3002'],
    allow_credentials=True,
    allow_methods=['GET', 'POST', 'PUT', 'DELETE', 'OPTIONS'],
    allow_headers=['*']
)

app.include_router(user_router, prefix='/users')
app.include_router(service_router, prefix='/services')
app.include_router(socket_router, prefix='/sockets')

async def send_one():
    producer = aiokafka.AIOKafkaProducer(
        bootstrap_servers='kafka:9092')
    # Get cluster layout and initial topic/partition leadership information
    await producer.start()
    try:
        # Produce message
        await producer.send_and_wait("main", b"Super message")
        print('message produced')
    finally:
        # Wait for all pending messages to be delivered or expire.
        
        await producer.stop()


@app.get("/")
async def root():
    await send_one()
    return {'message': 'pong'}


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
        'services': [
            {
                'name': 'Ressearch',
                'response': ressearch_response
            },
            {
                'name': 'Redis',
                'response': redis_response
            },
            {
                'name': 'Finance',
                'response': finance_response
            }
        ],
    }

