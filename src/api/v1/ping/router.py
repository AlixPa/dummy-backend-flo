from fastapi import APIRouter
from .schema import PingResponse
from .service import pong

router = APIRouter(prefix="/ping")


@router.get("", response_model=PingResponse)
async def ping():
    return PingResponse(message=pong())
