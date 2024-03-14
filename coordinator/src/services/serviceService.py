from repositories.serviceRepository import ServiceRepository
from fastapi import UploadFile
from schemas.serviceSchema import ServiceLogSchema, ServiceSchema, ServiceWithLogsSchema, \
      ExtendedServiceSchema, FullServiceSchema
from schemas.userSchema import UserSchema
from schemas.statusSchemas import MessageResponseSchema
from repositories.userRepository import UserRepository
import aiohttp
from aiohttp import FormData

class ServiceService:
    service_repository: ServiceRepository
    user_repository: UserRepository

    def __init__(self) -> None:
        self.service_repository = ServiceRepository()
        self.user_repository = UserRepository()
    
    async def create_service(self, username: str, servicedata: ServiceSchema) -> ServiceSchema:
        user: UserSchema = self.user_repository.get_user_by_name(username)
        
        if not user.is_staff:
            raise Exception('Not enough privileges!')
        
        service = self.service_repository.create_service(servicedata)
        return service
    
    async def upload_service_image(self, username: str, service_id: int, image: UploadFile):
        user: UserSchema = self.user_repository.get_user_by_name(username)
        
        if not user.is_staff:
            raise Exception('Not enough privileges!')
        
        service: FullServiceSchema = self.service_repository.get_service_by_id(service_id)
        data = FormData()
        data.add_field('avatar', image.file.read(), filename=image.filename, content_type=image.content_type)
        async with aiohttp.ClientSession() as session:
            async with session.post('http://support:3003/profiles/', data=data) as resp:
                data = await resp.json()
        print(data)
        self.service_repository.add_service_image(service.id, data['avatar'])
        return service


    
    async def get_services(self) -> list[ServiceSchema]:
        services: list[ServiceSchema] = self.service_repository.get_services()
        return services
    
    async def delete_service(self, service_id: int, username: str) -> MessageResponseSchema:
        user: UserSchema = self.user_repository.get_user_by_name(username)
        
        if not user.is_staff:
            raise Exception('Not enough privileges!')
        
        deleted = self.service_repository.delete_service(service_id)
        if deleted:
            return MessageResponseSchema(msg='Deleted')
        else:
            return MessageResponseSchema(msg='Failed to delete')
