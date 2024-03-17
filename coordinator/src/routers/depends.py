from services.userService import UserService
from services.serviceService import ServiceService
from fastapi import Request, HTTPException
from fastapi.security import HTTPBearer, HTTPAuthorizationCredentials

from utils.jwt_utils import decodeJWT
from typing import Tuple
from datetime import datetime


class JWTCredentials(HTTPAuthorizationCredentials):
    username: str


class JWTBearer(HTTPBearer):
    def __init__(self, auto_error: bool = True):
        super(JWTBearer, self).__init__(auto_error=auto_error)

    async def __call__(self, request: Request) -> JWTCredentials:
        credentials: HTTPAuthorizationCredentials = await super(JWTBearer, self).__call__(request)
        if credentials:
            if not credentials.scheme == "Bearer":
                raise HTTPException(status_code=403, detail="Invalid authentication scheme.")
            
            is_verified, username = self.verify_jwt(credentials.credentials)
            if not is_verified:
                raise HTTPException(status_code=403, detail="Invalid token or expired token.")
            
            
            return JWTCredentials(
                credentials=credentials.credentials,
                scheme=credentials.scheme,
                username=username
            )
        else:
            raise HTTPException(status_code=403, detail="Invalid authorization code.")

    def verify_jwt(self, jwtoken: str) -> Tuple[bool, str | None]:
        
        try:
            payload = decodeJWT(jwtoken)
        except:
            payload = None
        
        if payload:
            expires_at: str = payload.get('expires_at', None)
            
            if expires_at == None:
                return False, None
            
            if datetime.strptime(expires_at, '%Y-%m-%d %H:%M:%S.%f') < datetime.now():
                return False, None
            
            return True, payload.get('username')
        else:
            return False, None

jwtBearer = JWTBearer()

def get_jwt_bearer():
    return jwtBearer

user_service = UserService()
service_service = ServiceService()

def get_user_service():
    return user_service

def get_service_service():
    return service_service