from fastapi import Request, HTTPException
from fastapi.security import HTTPBearer, HTTPAuthorizationCredentials

from utils.jwt_utils import decodeJWT
from typing import Tuple

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
        isTokenValid: bool = False

        try:
            payload = decodeJWT(jwtoken)
        except:
            payload = None
        
        if payload:
            return True, payload.get('username')
        else:
            return False, None

jwtBearer = JWTBearer()

def get_jwt_bearer():
    return jwtBearer()