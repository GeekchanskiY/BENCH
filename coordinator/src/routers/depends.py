from services.userService import UserService
from services.serviceService import ServiceService

user_service = UserService()
service_service = ServiceService()

def get_user_service():
    return user_service

def get_service_service():
    return service_service