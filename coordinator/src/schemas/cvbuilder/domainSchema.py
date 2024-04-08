from pydantic import BaseModel, ConfigDict


class DomainSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)
    name: str

class DomainDetailedSchema(DomainSchema):
    id: int