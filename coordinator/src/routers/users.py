from fastapi import APIRouter, Depends, HTTPException, Request

from repositories.models.userModel import User
from repositories.models.db import get_db
from .depends import UserService, get_user_service
import schemas.userSchema as schemas
from sqlalchemy.orm import Session
from services.utils import bcrypt_utils

import jwt
from datetime import datetime, timedelta



router: APIRouter = APIRouter()

@router.get("/users/", tags=["users"], response_model=list[schemas.UserPrivateSchema])
async def read_users(db: Session = Depends(get_db)):
    users = db.query(User).all()
    return users

@router.post("/users/login", tags=["users"], response_model=schemas.UserSchema)
async def login_user(user: schemas.LoginUserSchema, service: UserService = Depends(get_user_service)):
    try:
        return await service.login(user)
    except Exception as e:
        raise HTTPException(404, str(e))

@router.get("/users/find/{userid}", tags=["users"])
async def read_user(username: str):
    return {"username": username}

@router.post("/users/register", tags=["users"])
async def create_user(user: schemas.RegisterUserSchema, request:Request, service: UserService = Depends(get_user_service)):
    try:
        return await service.register(user, request.client.host)
    except Exception as e:
        raise HTTPException(404, str(e))

@router.get("/users/delete", tags=["users"])
async def delete_user(userid: int, db: Session = Depends(get_db)):
    user = db.query(User).where(User.id == userid).first()
    if user is not None:
        db.delete(user)
        db.commit()
    
        return {'deleted': userid}
    
    raise HTTPException(status_code=404, detail="User not found")