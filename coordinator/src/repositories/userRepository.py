from typing import List

from sqlalchemy.orm import Session

from .models.db import get_db
from .models.userModel import User
from schemas.userSchema import UserSchema, RegisterUserSchema

class UserRepository:
    db: Session
    def __init__(self) -> None:
        self.db = get_db()

    def get_users(self) -> List[User]:
        return self.db.query(User).all()

    def get_user_by_email(self, email:str) -> User:
        return self.db.query(User).where(User.email == email).first()

    def get_user_by_name(self, name: str) -> User:
        return self.db.query(User).where(User.name == name).first()

    def create_user(self, user: RegisterUserSchema, ip: str, is_staff: bool) -> User:
        new_user = User()
        new_user.name = user.name
        new_user.email = user.email
        new_user.is_staff = is_staff
        new_user.ip_adress = ip
        new_user.password = user.password
        self.db.add(new_user)
        self.db.commit()
        return new_user
