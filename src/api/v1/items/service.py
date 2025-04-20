from .models import Item

items: dict[int, Item] = dict()


def get_items(limit: int) -> list[Item]:
    return [items[id] for id in items][:limit]


def add_item(item: Item) -> Item:
    new_id = len(items)
    items[new_id] = item
    return items[new_id]


def get_item_by_id(item_id: int) -> Item | None:
    if item_id in items:
        return items[item_id]
    else:
        return None
