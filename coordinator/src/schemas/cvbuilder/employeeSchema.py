from pydantic import BaseModel, ConfigDict


class EmployeeSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    age: int
    name: str

class EmployeeDetailedSchema(EmployeeSchema):
    id: int