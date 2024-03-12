from fastapi import APIRouter, WebSocket, WebSocketDisconnect
from utils.socketConnectionManager import ConnectionManager, get_connection_manager

router: APIRouter = APIRouter()

@router.websocket("/ws")
async def websocket_endpoint(websocket: WebSocket):
    manager: ConnectionManager = await get_connection_manager()
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