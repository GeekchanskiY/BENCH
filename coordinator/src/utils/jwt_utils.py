import jwt
from datetime import datetime, timedelta
from config import JWT_TOKEN_EXPIRE_MINUTES, JWT_SECRET


def decodeJWT(token) -> dict:
    return jwt.decode(token, JWT_SECRET, algorithms="HS256")

def encodeJWT(username: str) -> str:
    return jwt.encode({
            "username": username,
            "expires_at": str(
                datetime.now() + timedelta(minutes=JWT_TOKEN_EXPIRE_MINUTES)
                )
        }, JWT_SECRET, algorithm="HS256")