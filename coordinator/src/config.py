import os

JWT_TOKEN_EXPIRE_MINUTES: int = 30

JWT_SECRET = os.getenv('JWT_SECRET')