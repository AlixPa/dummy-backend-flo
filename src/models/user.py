from uuid import uuid4

from sqlalchemy import Column, Integer, String

from src._exceptions import WrongAttributesException
from src._models import BaseModel


class User(BaseModel):
    __tablename__ = "user"

    id = Column(String(38), primary_key=True)
    name = Column(String(50), nullable=False)
    age = Column(Integer, nullable=True)

    def __init__(self, attr: dict):
        self.id = attr.get("id") if attr.get("id") else str(uuid4())
        self.name = attr.get("name")
        self.age = attr.get("age")

    def check_validity_name(self):
        if isinstance(self.name, str) and (len(self.name) < 2 or len(self.name) > 50):
            raise WrongAttributesException(
                table_name=self.__tablename__,
                detail="name lenght must be between 2 and 50",
            )

    def check_validity_age(self):
        if isinstance(self.age, int) and (self.age < 0 or self.age > 120):
            raise WrongAttributesException(
                table_name=self.__tablename__, detail="age must be between 0 and 120"
            )

    def check_validity(self):
        self.check_validity_name()
        self.check_validity_age()

    def check_complete(self, skip_nullable: bool = False):
        if not isinstance(self.name, str):
            raise WrongAttributesException(
                table_name=self.__tablename__, detail="name must be present"
            )
        if not skip_nullable and not isinstance(self.age, int):
            raise WrongAttributesException(
                table_name=self.__tablename__, detail="age must be present"
            )
