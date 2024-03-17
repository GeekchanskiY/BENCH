from sqlalchemy import create_engine, MetaData, Column, String, Integer
from sqlalchemy.orm import Session
from sqlalchemy.orm import declarative_base
from sqlalchemy.orm import sessionmaker
import aioredis
from aioredis import Redis

from config import PG_DB, PG_HOST, PG_PORT, PG_PASSWORD, PG_USER

SQLALCHEMY_DATABASE_URL = f'postgresql://{PG_USER}:{PG_PASSWORD}@{PG_HOST}:{PG_PORT}/{PG_DB}'

engine = create_engine(SQLALCHEMY_DATABASE_URL)

SessionLocal: sessionmaker = sessionmaker(
    autocommit=False,
    autoflush=False, 
    bind=engine
)

session: Session = SessionLocal()

Base = declarative_base()

# Unused because of alembic
# Base.metadata.create_all(bind=engine)


async def get_redis() -> Redis:
    return await aioredis.from_url('redis://redis')


def get_db() -> Session:
    return session