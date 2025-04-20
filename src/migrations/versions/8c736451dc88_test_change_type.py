"""test change type

Revision ID: 8c736451dc88
Revises: 49783be5769e
Create Date: 2025-04-18 19:29:32.339235

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = '8c736451dc88'
down_revision: Union[str, None] = '49783be5769e'
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    pass


def downgrade() -> None:
    """Downgrade schema."""
    pass
