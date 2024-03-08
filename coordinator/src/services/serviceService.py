from repositories.serviceRepository import ServiceRepository
from schemas.serviceSchema import ServiceLogSchema, ServiceSchema, ServiceWithLogsSchema

class ServiceService:
    service_repository: ServiceRepository

    def __init__(self) -> None:
        self.service_repository = ServiceRepository()
    
    async def create_service(self, servicedata: ServiceSchema) -> ServiceSchema:
        service = self.service_repository.create_service(servicedata)
        return service
    
    async def get_services(self) -> list[ServiceSchema]:
        services: list[ServiceSchema] = self.service_repository.get_services()
        return services