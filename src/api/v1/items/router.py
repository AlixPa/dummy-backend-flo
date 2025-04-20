from fastapi import APIRouter, HTTPException
from .models import Item
from .schema import ListItemsResponse, GetItemResponse, CreateItemResponse
from .service import get_items, add_item, get_item_by_id


router = APIRouter(prefix="/items")


@router.post("", status_code=201)
def create_item(item: Item) -> CreateItemResponse:
    new_item = add_item(item=item)
    return CreateItemResponse(data=new_item)


@router.get("", response_model=ListItemsResponse)
def list_items(limit: int = 10) -> ListItemsResponse:
    items = get_items(limit=limit)
    return ListItemsResponse(data=items)


@router.get("/{item_id}", response_model=GetItemResponse)
def get_item(item_id: int) -> GetItemResponse:
    item = get_item_by_id(item_id=item_id)
    if item:
        return GetItemResponse(data=item)
    else:
        raise HTTPException(status_code=404, detail="Item not found")
