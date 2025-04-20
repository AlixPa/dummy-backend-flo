from fastapi import APIRouter
from .ping.router import router as ping_router
from .items.router import router as items_router
from .profiles.router import router as profiles_router

router = APIRouter(prefix="/v1")

router.include_router(ping_router)
router.include_router(items_router)
router.include_router(profiles_router)
