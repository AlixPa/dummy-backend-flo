from fastapi import APIRouter

from _schemas import DataResponse, MessageResponse
from src._database_pymysql import (
    NoConnectionError,
    NoUpdateValuesError,
    NoValueInsertionError,
)
from src._exceptions import (
    HTTPSNotFoundException,
    HTTPWrongAttributesException,
    NotFoundException,
    WrongAttributesException,
)

from .schema import ProfileCreateInput, ProfileUpdateInput
from .service import (
    add_profile,
    get_profile_by_id,
    get_profiles,
    remove_profile_by_id,
    update_profil_by_id,
)

router = APIRouter(prefix="/profiles")


@router.post("", status_code=201, response_model=MessageResponse)
def create_profile(profile_input: ProfileCreateInput) -> MessageResponse:
    try:
        profile = add_profile(profile_input=profile_input)
    except NoConnectionError as e:
        raise e
    except WrongAttributesException as e:
        raise HTTPWrongAttributesException(detail=str(e))
    return MessageResponse(message=f"Profile created with {profile.to_dict()}")


@router.get("", response_model=DataResponse[list[dict]])
def list_profiles(limit: int = 0, offset: int = 0) -> DataResponse[list[dict]]:
    try:
        profiles = get_profiles(limit=limit, offset=offset)
    except NoConnectionError as e:
        raise e
    return DataResponse(data=[p.to_dict() for p in profiles])


@router.get("/{id}", response_model=DataResponse[dict])
def get_profile(id: str) -> DataResponse[dict]:
    try:
        profile = get_profile_by_id(id=id)
    except NoConnectionError as e:
        raise e
    except NotFoundException as e:
        raise HTTPSNotFoundException(detail=str(e))
    return DataResponse(data=profile.to_dict())


@router.delete("/{id}", status_code=204)
def delete_profile(id: str):
    try:
        remove_profile_by_id(id=id)
    except NoConnectionError as e:
        raise e
    except NotFoundException as e:
        raise HTTPSNotFoundException(detail=str(e))


@router.put("/{id}", response_model=DataResponse[dict])
def put_profile(id: str, profile_input: ProfileUpdateInput):
    try:
        updated_profile = update_profil_by_id(id=id, profile_input=profile_input)
    except NotFoundException as e:
        raise HTTPSNotFoundException(detail=str(e))
    except WrongAttributesException as e:
        raise HTTPWrongAttributesException(detail=str(e))
    except NoConnectionError as e:
        raise e
    except NoUpdateValuesError as e:
        raise HTTPWrongAttributesException(detail=str(e))

    return DataResponse(data=updated_profile.to_dict())
