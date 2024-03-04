from fastapi import APIRouter

from models.user import User
from models.db import get_db


router = APIRouter()

@router.get("/users/", tags=["users"])
async def read_users():
    with get_db() as session:
        session.query(User).all()
    return [{"username": "Rick"}, {"username": "Morty"}]


@router.get("/users/me", tags=["users"])
async def read_user_me():
    return {"username": "fakecurrentuser"}


@router.get("/users/{username}", tags=["users"])
async def read_user(username: str):
    return {"username": username}