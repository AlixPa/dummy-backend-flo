"""test change type

Revision ID: def5506f3e85
Revises: 8c736451dc88
Create Date: 2025-04-18 19:32:37.940831

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = 'def5506f3e85'
down_revision: Union[str, None] = '8c736451dc88'
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    """Upgrade schema."""
    pass


def downgrade() -> None:
    """Downgrade schema."""
    pass
