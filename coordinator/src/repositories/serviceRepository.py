from typing import List

from sqlalchemy.orm import Session

from .models.db import get_db
from .models.serviceModel import Service

class ServiceRepository:
    db: Session

    def __init__(self):
        self.db = get_db()

    def get_services(self) -> List[Service]:
        pass

    def get_service_by_name(self, name: str) -> Service:
        pass

    def create_service(self) -> Service:
        pass

    def delete_service(self) -> bool:
        pass