from .db import Base

from sqlalchemy import Column, String, Integer, Boolean

class User(Base):
     __tablename__ = "user_account"

     id: Column = Column(Integer, primary_key=True)
     name: Column = Column(String(30))
     email: Column = Column(String)
     password: Column = Column(String)
     is_staff: Column = Column(Boolean)
     ip_adress: Column = Column(String)

     def __repr__(self) -> str:
         return f"User(id={self.id!r}, name={self.name!r}, fullname={self.fullname!r})"
