from fastapi import APIRouter, Depends, HTTPException, Request

from .depends import get_service_service
from services.serviceService import ServiceService
from schemas.serviceSchema import ServiceSchema, ServiceLogSchema, ServiceWithLogsSchema

router: APIRouter = APIRouter()

@router.get("/", tags=["services"], response_model=list[ServiceSchema])
async def read_services(service: ServiceService = Depends(get_service_service)):
    services = await service.get_services()
    return services

@router.post('/create', tags=['services'], response_model=ServiceSchema)
async def create_service(service_data: ServiceSchema, service: ServiceService = Depends(get_service_service)):
    created_service = await service.create_service(service_data)
    return created_service