from typing import List

from sqlalchemy.orm import Session

from .models.db import get_db
from .models.userModel import User
from .schemas.userSchema import UserSchema, UserPrivateSchema

class UserRepository:
    db: Session
    def __init__(self) -> None:
        self.db = get_db()

    def get_users(self) -> List[User]:
        return self.db.query(User).all()

    def get_user_by_email(self, email:str) -> User:
        pass

    def get_user_by_name(self, name: str) -> User:
        pass

    def create_user(self, user: UserPrivateSchema) -> User:
        pass
