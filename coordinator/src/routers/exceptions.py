from typing import Any, Dict
from typing_extensions import Annotated, Doc
from fastapi import HTTPException

class NotAuthentificatedException(HTTPException):
    def __init__(
            self,
            status_code: int = 401,
            detail: Any = 'Unauthorized',
            headers: Dict[str, str] | None = None
        ) -> None:
        super().__init__(status_code, detail, headers)

async def exceptionHandler(exception: Exception):
    return HTTPException(400, str(exception))