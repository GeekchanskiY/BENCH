from .db import Base

from typing import List

from sqlalchemy import Column, String, Integer, Boolean, ForeignKey
from sqlalchemy.orm import relationship, Mapped, mapped_column

class Service(Base):
     __tablename__ = "servises"

     id: Mapped[int] = mapped_column(primary_key=True)
     name: Mapped[String] = Column(String(30))
     description: Mapped[String] = Column(String)

     is_active: Column = Column(Boolean)

     logs: Mapped[List['ServiceLog']] = relationship(back_populates="servise")
     

     def __repr__(self) -> str:
         return f"User(id={self.id!r}, name={self.name!r}, fullname={self.fullname!r})"

class ServiceLog(Base):
    __tablename__ = "serviselogs"
    id: Mapped[int] = mapped_column(primary_key=True)
    service_id: Mapped[int] = mapped_column(ForeignKey(Service.id))
    service: Mapped['Service'] = relationship(back_populates='service_logs')