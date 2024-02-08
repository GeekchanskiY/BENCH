from .db import Base

from sqlalchemy import Column, String, Integer

class User(Base):
     __tablename__ = "user_account"

     id: Column = Column(Integer, primary_key=True)
     name: Column = Column(String(30))


     def __repr__(self) -> str:
         return f"User(id={self.id!r}, name={self.name!r}, fullname={self.fullname!r})"
