from src._database_pymysql import (
    MysqlClient,
    NoConnectionError,
    NoUpdateValuesError,
    NoValueInsertionError,
)
from src._exceptions import NotFoundException, WrongAttributesException
from src.models import User

from .schema import ProfileCreateInput, ProfileUpdateInput


def get_profiles(limit: int = 0, offset: int = 0) -> list[User]:
    mysql_client = MysqlClient()
    try:
        res_sql = mysql_client.select(
            table_name=User.__tablename__, limit=limit, offset=offset
        )
    except NoConnectionError as e:
        raise e
    mysql_client.close_connection()
    ls_res = [User(p) for p in res_sql]
    return ls_res


def add_profile(profile_input: ProfileCreateInput) -> User:
    profile = User(attr=profile_input.model_dump(exclude_unset=True))
    try:
        profile.check_complete(skip_nullable=True)
        profile.check_validity()
    except WrongAttributesException as e:
        raise e

    mysql_client = MysqlClient()
    try:
        mysql_client.insert_one(
            table_name=profile.__tablename__, values=profile.to_dict()
        )
    except NoConnectionError as e:
        raise e
    except NoValueInsertionError as e:
        raise e
    mysql_client.close_connection()

    return profile


def get_profile_by_id(id: str) -> User:
    mysql_client = MysqlClient()
    try:
        res_mysql = mysql_client.select_by_id(table_name=User.__tablename__, id=id)
    except NoConnectionError as e:
        raise e
    mysql_client.close_connection()
    if not res_mysql:
        raise NotFoundException(table_name=User.__tablename__, detail=f"{id=}")
    return User(attr=res_mysql)


def remove_profile_by_id(id: str):
    mysql_client = MysqlClient()
    try:
        res_mysql = mysql_client.delete_by_id(table_name=User.__tablename__, id=id)
    except NoConnectionError as e:
        raise e
    mysql_client.close_connection()
    if not res_mysql:
        raise NotFoundException(table_name=User.__tablename__, detail=f"{id=}")


def update_profil_by_id(id: str, profile_input: ProfileUpdateInput) -> User:
    profile = User(attr=profile_input.model_dump(exclude_unset=True))
    try:
        profile.check_validity()
    except WrongAttributesException as e:
        raise e

    mysql_client = MysqlClient()

    try:
        mysql_res = mysql_client.update_by_id(
            table_name=User.__tablename__,
            id=id,
            values=profile.to_dict(exclude_null=True),
        )
    except NoConnectionError as e:
        raise e
    except NoUpdateValuesError as e:
        raise e

    mysql_client.close_connection()

    if not mysql_res:
        raise NotFoundException(table_name=User.__tablename__, detail=f"{id=}")

    return User(attr=mysql_res)
