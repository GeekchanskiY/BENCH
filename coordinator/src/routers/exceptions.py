from fastapi import HTTPException


async def exceptionHandler(exception: Exception):
    return HTTPException(400, 'Bad request')