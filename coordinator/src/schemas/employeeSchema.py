from pydantic import BaseModel, ConfigDict
from datetime import datetime

class EmployeeSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)
    name: str
    age: int

class EmployeeFullSchema(EmployeeSchema):
    CreatedAt: datetime
    UpdatedAt: datetime
    ID: int