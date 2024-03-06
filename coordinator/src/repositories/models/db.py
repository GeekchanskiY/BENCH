from sqlalchemy import create_engine, MetaData, Column, String, Integer
from sqlalchemy.orm import Session
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker
import aioredis
from aioredis import Redis

SQLALCHEMY_DATABASE_URL = 'postgresql://staffing:staffing@db-fastapi:5432/backend_fastapi'

engine = create_engine(SQLALCHEMY_DATABASE_URL)
SessionLocal: sessionmaker = sessionmaker(autocommit=False, autoflush=False, bind=engine)
session: Session = SessionLocal()

Base = declarative_base()
Base.metadata.create_all(bind=engine)


async def get_redis() -> Redis:
    return await aioredis.from_url('redis://redis')


def get_db() -> Session:
    return session