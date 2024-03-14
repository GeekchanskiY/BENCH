from typing import List

from sqlalchemy.orm import Session

from .models.db import get_db
from .models.serviceModel import Service, ServiceLog
from schemas.serviceSchema import ServiceSchema, ServiceLogSchema

class ServiceRepository:
    db: Session

    def __init__(self):
        self.db = get_db()

    def get_services(self) -> List[Service]:
        return self.db.query(Service).all()

    def get_service_by_id(self, service_id: int) -> Service:
        return self.db.query(Service).where(Service.id == service_id).first()

    def get_service_by_name(self, name: str) -> Service:
        return self.db.query(Service).where(Service.name == name).first()

    def create_service(self, service: ServiceSchema) -> Service:
        new_service: Service = Service()
        new_service.description = service.description
        new_service.name = service.name
        new_service.is_active = service.is_active
        new_service.url = service.url
        new_service.ping_url = service.ping_url
        new_service.image_url = service.image_url

        try:
            self.db.add(new_service)
            self.db.commit()
            return new_service
        except Exception as e:
            self.db.rollback()
            raise e

    def delete_service(self, service_id: int) -> bool:
        service: Service = self.db.query(Service).where(Service.id == service_id).first()
        if service is None:
            raise Exception('Service not found')
        
        try:
            self.db.delete(service)
            self.db.commit()
            return True
        except Exception as e:
            self.db.rollback()
            raise e

    def create_service_log(self, servicelog: ServiceLogSchema) -> ServiceLog:
        new_servicelog: ServiceLog = ServiceLog()
        service: Service = self.db.query(Service).where(Service.id == servicelog.service_id).first()
        if service is None:
            raise Exception('Related service not found')
        
        new_servicelog.service = service
        new_servicelog.service_id = service.id
        new_servicelog.message = servicelog.message

        try:
            self.db.add(new_servicelog)
            self.db.commit()
            return new_servicelog
        except Exception as e:
            self.db.rollback()
            raise e

    def delete_service_log(self, servicelog_id: int) -> bool:
        servicelog: ServiceLog = self.db.query(ServiceLog).where(ServiceLog.id == servicelog_id).first()
        if servicelog is None:
            raise Exception('Service log not found!')
        try:
            self.db.delete(servicelog)
            self.db.commit()
        except Exception as e:
            self.db.rollback()
            raise e
 