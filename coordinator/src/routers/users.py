from fastapi import APIRouter, Depends, HTTPException

from models.user import User
from models.db import get_db
import schemas.user as schemas
from sqlalchemy.orm import Session



router: APIRouter = APIRouter()

@router.get("/users/", tags=["users"], response_model=list[schemas.User])
async def read_users(db: Session = Depends(get_db)):
    users = db.query(User).all()
    return users


@router.get("/users/me", tags=["users"])
async def read_user_me():
    return {"username": "fakecurrentuser"}


@router.get("/users/{username}", tags=["users"])
async def read_user(username: str):
    return {"username": username}

@router.post("/users/create", tags=["users"])
async def create_user(user: schemas.RegisterUser, db: Session = Depends(get_db)):
    new_user = User()
    new_user.name = user.name
    new_user.email = user.email
    new_user.is_staff = True
    new_user.ip_adress = 'localhost'
    db.add(new_user)
    db.commit()
    return new_user

@router.get("/users/{userid}/delete")
async def delete_user(userid: int, db: Session = Depends(get_db)):
    user = db.query(User).where(User.id == userid).first()
    if user is not None:
        db.delete(user)
        db.commit()
    
        return {'deleted': userid}
    
    raise HTTPException(status_code=404, detail="User not found")