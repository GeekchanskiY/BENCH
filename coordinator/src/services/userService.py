from repositories.userRepository import UserRepository

from schemas.userSchema import UserSchema, LoginUserSchema, UserPrivateSchema \
    , RegisterUserSchema
from schemas.jwtSchema import JWTDetailedSchema, JWTSchema
from schemas.statusSchemas import MessageResponseSchema

from utils import bcrypt_utils, jwt_utils


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
        
        token: str
        expires_at: str
        token, expires_at = jwt_utils.encodeJWT(userdata.email)

        return JWTDetailedSchema(
            token=token,
            username=db_user.name,
            expires_at=expires_at
        )

    async def refresh_login(self, token: JWTSchema) -> JWTDetailedSchema:
        token_data = jwt_utils.decodeJWT(token.token)
        
        return {
            'asd'
        }

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
    

    async def get_user_by_name(self, username: str) -> UserSchema:

        user = self.user_repository.get_user_by_name(username)
        if user is not None:
            return UserSchema(
                id=user.id,
                ip_adress=user.ip_adress,
                name=user.name,
                email=user.email,
                is_staff=user.is_staff,
            )
        else:
            raise Exception('User not found')
        
    async def get_users(self) -> list[UserSchema]:
        return self.user_repository.get_users()
    
    async def delete_user(self, username: str) -> bool:
        deleted: bool = self.user_repository.delete_user(username)
        if deleted:
            return MessageResponseSchema(msg='Deleted')
        else:
            return MessageResponseSchema(msg='Failed to delete')
    