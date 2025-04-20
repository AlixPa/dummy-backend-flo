from .models import Item
from pydantic import BaseModel


class ListItemsResponse(BaseModel):
    data: list[Item]


class GetItemResponse(BaseModel):
    data: Item


class CreateItemResponse(BaseModel):
    data: Item
