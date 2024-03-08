"""

Revision ID: 0bb916ab5e4b
Revises: 217033334431
Create Date: 2024-03-08 02:03:24.993446

"""
from typing import Sequence, Union

from alembic import op
import sqlalchemy as sa


# revision identifiers, used by Alembic.
revision: str = '0bb916ab5e4b'
down_revision: Union[str, None] = '217033334431'
branch_labels: Union[str, Sequence[str], None] = None
depends_on: Union[str, Sequence[str], None] = None


def upgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.add_column('servicelogs', sa.Column('message', sa.String(), nullable=False))
    # ### end Alembic commands ###


def downgrade() -> None:
    # ### commands auto generated by Alembic - please adjust! ###
    op.drop_column('servicelogs', 'message')
    # ### end Alembic commands ###
