from fastapi import APIRouter, Depends, HTTPException, Request

from .depends import get_service_service
from services.serviceService import ServiceService
from schemas.serviceSchema import ServiceSchema, ServiceLogSchema, ServiceWithLogsSchema
from schemas.statusSchemas import MessageResponseSchema

from .depends import JWTBearer, get_jwt_bearer, JWTCredentials


router: APIRouter = APIRouter()

@router.get("/", tags=["services"], response_model=list[ServiceSchema])
async def read_services(
    service: ServiceService = Depends(get_service_service),
    ):
    services = await service.get_services()
    return services

@router.post('/create', tags=['services'], response_model=ServiceSchema)
async def create_service(
    service_data: ServiceSchema,
    service: ServiceService = Depends(get_service_service),
    credentials: JWTCredentials = Depends(JWTBearer())):

    created_service = await service.create_service(credentials.username, service_data)
    return created_service

@router.delete('/service_id/{service_id}', tags=['services'], response_model=MessageResponseSchema)
async def delete_service(
    service_id: int,
    service: ServiceService = Depends(get_service_service),
    credentials: JWTCredentials = Depends(JWTBearer())):
    
    try:
        return await service.delete_service(service_id, credentials.username)
    except Exception as e:
        raise HTTPException(404, str(e))
    