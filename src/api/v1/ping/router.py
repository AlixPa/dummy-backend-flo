from fastapi import APIRouter

from src._schemas import MessageResponse

from .service import pong

router = APIRouter(prefix="/ping")


@router.get("", response_model=MessageResponse)
async def ping():
    return MessageResponse(message=pong())
