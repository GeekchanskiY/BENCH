import aiohttp

from fastapi import FastAPI, WebSocket, WebSocketDisconnect

from routers.users import router as user_router
from routers.services import router as service_router

from fastapi.middleware.cors import CORSMiddleware

from repositories.models.db import get_db, get_redis
import websockets
import pika
import asyncio

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=['*'],
    allow_credentials=True,
    allow_methods=['*'],
    allow_headers=['*']
)

app.include_router(user_router, prefix='/users')
app.include_router(service_router, prefix='/services')

@app.get("/")
async def root():
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

class ConnectionManager:
    def __init__(self):
        self.active_connections: list[WebSocket] = []

    async def connect(self, websocket: WebSocket):
        await websocket.accept()
        self.active_connections.append(websocket)

    def disconnect(self, websocket: WebSocket):
        self.active_connections.remove(websocket)

    async def send_personal_message(self, message: str, websocket: WebSocket):
        await websocket.send_text(message)

    async def broadcast(self, message: str):
        for connection in self.active_connections:
            await connection.send_text(message)


manager: ConnectionManager = ConnectionManager()

@app.websocket("/ws")
async def websocket_endpoint(websocket: WebSocket):
    await manager.connect(websocket)
    await manager.broadcast(f"Current clients: {len(manager.active_connections)}")
    try:
        while True:
           
            recieved_message = await websocket.receive_text()
            match recieved_message:
                case 'On my way!':
                    await websocket.send_json(
                        {
                            'message': f'Hi there! Current connections: {len(manager.active_connections)}'
                        }
                    )
                
                case 'Close connection':
                    manager.disconnect(websocket)
                    await websocket.close()
                    break
                case _:
                    await websocket.send_json(f'Message recieved: {recieved_message}' )
            
    except WebSocketDisconnect:
        manager.disconnect(websocket)
        await manager.broadcast(f"Client #{client_id} left the chat..")
    except Exception as e:
        print(str(e))
        await manager.broadcast(f"{str(e)}")
        manager.disconnect(websocket)
        