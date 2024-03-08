from fastapi import APIRouter, Depends, HTTPException, Request

from repositories.models.userModel import User
from repositories.models.db import get_db
from .depends import UserService, get_user_service
from schemas.userSchema import UserSchema, LoginUserSchema, UserPrivateSchema, RegisterUserSchema
from schemas.jwtSchema import JWTDetailedSchema
from sqlalchemy.orm import Session

from .middlewares import JWTBearer



router: APIRouter = APIRouter()

@router.get("/", tags=["users"], response_model=list[UserPrivateSchema])
async def read_users(db: Session = Depends(get_db)):
    users = db.query(User).all()
    return users

@router.post("/auth", tags=["users"], response_model=JWTDetailedSchema)
async def authenticate_user(user: LoginUserSchema, service: UserService = Depends(get_user_service)):
    try:
        return await service.login(user)
    except Exception as e:
        raise HTTPException(404, str(e))

@router.get("/find/{userid}", dependencies=[Depends(JWTBearer())], tags=["users"])
async def read_user(username: str, service: UserService = Depends(get_user_service)):
    return {"username": username}

@router.get("/whoami", dependencies=[Depends(JWTBearer())], tags=["users"])
async def whoami(request:Request, service: UserService = Depends(get_user_service)):
    try:
        return await service.whoami(token=request.headers.get('Authorization'))
    except Exception as e:
        raise HTTPException(404, str(e))

@router.post("/register", tags=["users"])
async def create_user(user: RegisterUserSchema, request:Request, service: UserService = Depends(get_user_service)):
    try:
        return await service.register(user, request.client.host)
    except Exception as e:
        raise HTTPException(404, str(e))

@router.get("/delete", tags=["users"])
async def delete_user(userid: int, db: Session = Depends(get_db)):
    user = db.query(User).where(User.id == userid).first()
    if user is not None:
        db.delete(user)
        db.commit()
    
        return {'deleted': userid}
    
    raise HTTPException(status_code=404, detail="User not found")