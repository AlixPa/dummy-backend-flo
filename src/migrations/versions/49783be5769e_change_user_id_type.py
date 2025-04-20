"""change user id type

Revision ID: 49783be5769e
Revises: d7f149ec1144
Create Date: 2025-04-18 19:24:17.078118

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = '49783be5769e'
down_revision: Union[str, None] = 'd7f149ec1144'
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    pass


def downgrade() -> None:
    """Downgrade schema."""
    pass
