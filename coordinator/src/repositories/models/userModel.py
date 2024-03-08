from .db import Base

from sqlalchemy import Column, String, Integer, Boolean


class User(Base):
     __tablename__ = "users"

     id: Column = Column(Integer, primary_key=True)
     name: Column = Column(String(30), unique=True)
     email: Column = Column(String, unique=True)
     password: Column = Column(String)
     is_staff: Column = Column(Boolean)
     ip_adress: Column = Column(String)

     def __repr__(self) -> str:
         return f"User(id={self.id!r}," \
         "name={self.name!r}, fullname={self.fullname!r})"
