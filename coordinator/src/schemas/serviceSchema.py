from pydantic import BaseModel, ConfigDict, validator

class ServiceSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    id: int | None

    name: str
    description: str
    is_active: bool

    url: str | None
    ping_url: str | None


class ServiceLogSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    service_id: int
    message: str

class ServiceWithLogsSchema(ServiceSchema):
    logs: list[ServiceLogSchema]
