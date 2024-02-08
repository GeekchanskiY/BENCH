from sqlalchemy import create_engine, MetaData, Column, String, Integer
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
import aioredis

SQLALCHEMY_DATABASE_URL = 'postgresql://staffing:staffing@db-fastapi:5432/backend_fastapi'

engine = create_engine(SQLALCHEMY_DATABASE_URL)
metadata_obj = MetaData()
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base()


class User(Base):
     __tablename__ = "user_account"

     id: Column = Column(Integer, primary_key=True)
     name: Column = Column(String(30))


     def __repr__(self) -> str:
         return f"User(id={self.id!r}, name={self.name!r}, fullname={self.fullname!r})"


async def get_redis():
    redis = await aioredis.from_url('redis://redis')
    return redis
    

def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()