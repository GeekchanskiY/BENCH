from fastapi import APIRouter, Depends, HTTPException, Request, UploadFile, File

from .depends import get_service_service
from services.serviceService import ServiceService
from schemas.serviceSchema import ServiceSchema, ServiceLogSchema, ServiceWithLogsSchema, \
      ExtendedServiceSchema, FullServiceSchema
from schemas.statusSchemas import MessageResponseSchema

from .depends import get_jwt_bearer, JWTCredentials


router: APIRouter = APIRouter()

@router.get("/", tags=["services"], response_model=list[FullServiceSchema])
async def read_services(
    service: ServiceService = Depends(get_service_service),
    ):
    services = await service.get_services()
    return services

@router.post('/create', tags=['services'], response_model=ServiceSchema)
async def create_service(
    service_data: ServiceSchema,
    service: ServiceService = Depends(get_service_service),
    credentials: JWTCredentials = Depends(get_jwt_bearer())):

    created_service = await service.create_service(credentials.username, service_data)
    return created_service

@router.delete('/service/{service_id}', tags=['services'], response_model=MessageResponseSchema)
async def delete_service(
    service_id: int,
    service: ServiceService = Depends(get_service_service),
    credentials: JWTCredentials = Depends(get_jwt_bearer())):
    
    try:
        return await service.delete_service(service_id, credentials.username)
    except Exception as e:
        raise HTTPException(404, str(e))

@router.post('/service/{service_id}', tags=['services'], response_model=ServiceSchema)
async def upload_service_image(
    service_id: int,
    image: UploadFile,
    service: ServiceService = Depends(get_service_service),
    credentials: JWTCredentials = Depends(get_jwt_bearer())):

    try:
        return await service.upload_service_image(credentials.username, service_id, image)
    except Exception as e:
        raise HTTPException(404, str(e))


    