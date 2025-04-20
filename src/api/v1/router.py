from fastapi import APIRouter

from .ping import ping_router
from .profiles import profiles_router

router = APIRouter(prefix="/v1")

router.include_router(ping_router)
router.include_router(profiles_router)
