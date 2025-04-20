from pydantic import BaseModel


class ProfileCreateInput(BaseModel):
    name: str | None = None
    age: int | None = None


class ProfileUpdateInput(BaseModel):
    name: str | None = None
    age: int | None = None
