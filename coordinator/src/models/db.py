from sqlalchemy import create_engine, MetaData, Column, String, Integer
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
import aioredis

SQLALCHEMY_DATABASE_URL = 'postgresql://staffing:staffing@db-fastapi:5432/backend_fastapi'

engine = create_engine(SQLALCHEMY_DATABASE_URL)
SessionLocal = sessionmaker(autocommit=False, autoflush=False, bind=engine)

Base = declarative_base()

async def get_redis():
    redis = await aioredis.from_url('redis://redis')
    return redis
    

def get_db():
    Base.metadata.create_table()
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()