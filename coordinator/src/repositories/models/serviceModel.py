from .db import Base

from typing import List

from sqlalchemy import Column, String, Integer, Boolean, ForeignKey
from sqlalchemy.orm import relationship, Mapped, mapped_column

class Service(Base):
    __tablename__ = "services"

    id: Mapped[int] = mapped_column(primary_key=True)
    name: Mapped[String] = Column(String(30), unique=True)
    description: Mapped[String] = Column(String)
    is_active: Column = Column(Boolean)

    url: Column = Column(String)
    ping_url: Column = Column(String)

    logs: Mapped[List['ServiceLog']] = relationship(
        'ServiceLog',
        back_populates='service'
    )
     
    def __repr__(self) -> str:
        return f"Service(id={self.id!r}," \
            "name={self.name!r}," \
            "description={self.description!r})"

class ServiceLog(Base):
    __tablename__ = "servicelogs"
    id: Mapped[int] = mapped_column(primary_key=True)
    service_id: Mapped[int] = mapped_column(
        ForeignKey('services.id'),
        nullable=False
    )
    message: Mapped[String] = Column(String, nullable=False)
    service: Mapped[Service] = relationship('Service', back_populates='logs')