import jwt
from datetime import datetime, timedelta
from repositories.userRepository import UserRepository
from schemas.userSchema import UserSchema, LoginUserSchema, UserPrivateSchema \
    , RegisterUserSchema
from schemas.jwtSchema import JWTDetailedSchema
from utils import bcrypt_utils


class UserService:
    user_repository: UserRepository

    def __init__(self):
        self.user_repository = UserRepository()

    async def login(self, userdata: LoginUserSchema) -> JWTDetailedSchema:
        db_user: UserPrivateSchema = self.user_repository.get_user_by_email(userdata.email)
        if db_user is None:
            raise Exception("User not found!")
        
        if not bcrypt_utils.verify_password(db_user.password, userdata.password):
            raise Exception("Incorrect password!")

        token: str = jwt.encode({
            "username": db_user.name,
            "expires_at": str(datetime.now() + timedelta(minutes=30))
        }, "secret", algorithm="HS256")

        return JWTDetailedSchema(
            token=token,
            username=db_user.name,
            expires_at=str(datetime.now() + timedelta(minutes=30))
        )

    async def register(self, userdata: RegisterUserSchema, ip: str) -> UserSchema:
        userdata.password = bcrypt_utils.hash_password(userdata.password).decode('utf-8')
        user = self.user_repository.create_user(userdata, ip, False)
        return UserSchema(
            id=user.id,
            ip_adress=user.ip_adress,
            name=user.name,
            email=user.email,
            is_staff=user.is_staff,
        )
    