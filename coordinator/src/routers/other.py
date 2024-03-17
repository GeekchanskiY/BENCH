from fastapi import APIRouter, HTTPException
from fastapi.responses import Response
import aiohttp
from utils.fetch_image import fetch_image
from .exceptions import exceptionHandler


router: APIRouter = APIRouter()

@router.get('/image/{domain}/{image_name}',
            tags=['other'],
            responses = {
                200: {
                    "content": {"image/png": {}}
                }
            },
            response_class=Response)
async def get_image(domain: str, image_name: str):
    try:
        response = await fetch_image(f'http://support:3003/media/{domain}/{image_name}')
        return Response(content=response, media_type="image/jpg")
    except Exception as e:
        raise await exceptionHandler(e)