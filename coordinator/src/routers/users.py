from fastapi import APIRouter, Depends, HTTPException, Request
from fastapi.security import HTTPAuthorizationCredentials
from typing import Annotated
from repositories.models.userModel import User
from repositories.models.db import get_db
from .depends import UserService, get_user_service
from schemas.userSchema import UserSchema, LoginUserSchema, UserPrivateSchema, RegisterUserSchema
from schemas.jwtSchema import JWTDetailedSchema
from sqlalchemy.orm import Session
from schemas.statusSchemas import MessageResponseSchema
from .depends import JWTBearer, get_jwt_bearer, JWTCredentials



router: APIRouter = APIRouter()

@router.get("/", tags=["users"], response_model=list[UserPrivateSchema])
async def read_users(service: UserService = Depends(get_user_service)):
    try:
        return await service.get_users()
    except Exception as e:
        raise HTTPException(400, str(e))

@router.post("/auth", tags=["users"], response_model=JWTDetailedSchema)
async def authenticate_user(user: LoginUserSchema, service: UserService = Depends(get_user_service)):
    try:
        return await service.login(user)
    except Exception as e:
        raise HTTPException(400, str(e))

@router.get("/get/{userid}", dependencies=[Depends(JWTBearer())], tags=["users"])
async def read_user(username: str, service: UserService = Depends(get_user_service)):
    try:
        return await service.get_user_by_name(username)
    except Exception as e:
        raise HTTPException(400, str(e))

@router.get("/whoami", tags=["users"], response_model=UserSchema)
async def whoami(credentials: JWTCredentials = Depends(JWTBearer()), service: UserService = Depends(get_user_service),):
    try:
        return await service.get_user_by_name(username=credentials.username)
    except Exception as e:
        raise HTTPException(400, str(e))

@router.post("/register", tags=["users"])
async def create_user(user: RegisterUserSchema, request:Request, service: UserService = Depends(get_user_service)):
    try:
        return await service.register(user, request.client.host)
    except Exception as e:
        raise HTTPException(404, str(e))

@router.get("/user/{username}", tags=["users"], response_model=MessageResponseSchema)
async def delete_user(username: str, service: UserService = Depends(get_user_service)):
    try:
        return await service.delete_user(username)
    except Exception as e:
        raise HTTPException(status_code=404, detail=str(e))