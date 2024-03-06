from fastapi import APIRouter, Depends, HTTPException

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


@router.get("/users/me", tags=["users"])
async def read_user_me():
    return {"username": "fakecurrentuser"}

@router.post("/users/login", tags=["users"])
async def login_user(user: schemas.LoginUserSchema, service: UserService = Depends(get_user_service)):
    try:
        return await service.login(user)
    except Exception as e:
        raise HTTPException(404, str(e))

@router.post("/users/login/refresh", tags=["users"])
async def login_refresh():
    return {
        'status': 'success',
        'token': 'token',
        'expires_in': 9000
    }

@router.get("/users/{userid}", tags=["users"])
async def read_user(username: str):
    return {"username": username}

@router.post("/users/create", tags=["users"])
async def create_user(user: schemas.RegisterUser, db: Session = Depends(get_db)):
    new_user = User()
    new_user.name = user.name
    new_user.email = user.email
    new_user.is_staff = True
    new_user.ip_adress = 'localhost'
    new_user.password = bcrypt_utils.hash_password(user.password).decode('utf-8')
    db.add(new_user)
    db.commit()
    return new_user

@router.get("/users/{userid}/delete", tags=["users"])
async def delete_user(userid: int, db: Session = Depends(get_db)):
    user = db.query(User).where(User.id == userid).first()
    if user is not None:
        db.delete(user)
        db.commit()
    
        return {'deleted': userid}
    
    raise HTTPException(status_code=404, detail="User not found")