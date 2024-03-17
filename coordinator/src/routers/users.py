from fastapi import APIRouter, Depends, HTTPException, Request
from fastapi.security import HTTPAuthorizationCredentials
from typing import Annotated
from repositories.models.userModel import User
from repositories.models.db import get_db
from .depends import UserService, get_user_service
from schemas.userSchema import UserSchema, LoginUserSchema, UserPrivateSchema, RegisterUserSchema
from schemas.jwtSchema import JWTDetailedSchema, JWTSchema
from sqlalchemy.orm import Session
from schemas.statusSchemas import MessageResponseSchema
from .depends import JWTBearer, get_jwt_bearer, JWTCredentials
from .exceptions import exceptionHandler



router: APIRouter = APIRouter()

@router.get("/", tags=["users"], response_model=list[UserPrivateSchema])
async def read_users(service: UserService = Depends(get_user_service)):
    try:
        return await service.get_users()
    except Exception as e:
        raise await exceptionHandler(e)

@router.post("/auth", tags=["users"], response_model=JWTDetailedSchema)
async def authenticate_user(user: LoginUserSchema, service: UserService = Depends(get_user_service)):
    try:
        return await service.login(user)
    except Exception as e:
        raise await exceptionHandler(e)

@router.post("/auth/refresh", tags=["users", "jwt"], response_model=JWTDetailedSchema)
async def refresh_token(
    token: JWTSchema,
    service: UserService = Depends(get_user_service),
    _: JWTCredentials = Depends(get_jwt_bearer())):
    try:
        return await service.refresh_login(token)
    except Exception as e:
        raise await exceptionHandler(e)

@router.get("/get/{userid}", dependencies=[Depends(JWTBearer())], tags=["users"])
async def read_user(username: str, service: UserService = Depends(get_user_service)):
    try:
        return await service.get_user_by_name(username)
    except Exception as e:
        raise await exceptionHandler(e)

@router.get("/whoami", tags=["users"], response_model=UserSchema)
async def whoami(credentials: JWTCredentials = Depends(JWTBearer()), service: UserService = Depends(get_user_service),):
    try:
        return await service.get_user_by_name(username=credentials.username)
    except Exception as e:
        raise await exceptionHandler(e)

@router.post("/register", tags=["users"])
async def create_user(user: RegisterUserSchema, request:Request, service: UserService = Depends(get_user_service)):
    try:
        return await service.register(user, request.client.host)
    except Exception as e:
        raise await exceptionHandler(e)

@router.get("/user/{username}", tags=["users"], response_model=MessageResponseSchema)
async def delete_user(username: str, service: UserService = Depends(get_user_service)):
    try:
        return await service.delete_user(username)
    except Exception as e:
        raise await exceptionHandler(e)