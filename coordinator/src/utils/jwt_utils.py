import jwt
from datetime import datetime, timedelta
from config import JWT_TOKEN_EXPIRE_MINUTES, JWT_SECRET


def decodeJWT(token) -> dict:
    return jwt.decode(token, JWT_SECRET, algorithms="HS256")

def encodeJWT(username: str) -> tuple[str, str]:

    expires_at: str = str(
        datetime.now() + timedelta(minutes=JWT_TOKEN_EXPIRE_MINUTES)
    )

    return (
        jwt.encode({
                "username": username,
                "expires_at": expires_at
            },
            JWT_SECRET,
            algorithm="HS256"
        ),
        expires_at)