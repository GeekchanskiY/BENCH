from pydantic import BaseModel, ConfigDict

class ServiceSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    name: str
    description: str
    is_active: bool

    url: str | None
    ping_url: str | None

class ExtendedServiceSchema(ServiceSchema):
    image_url: str | None

class FullServiceSchema(ExtendedServiceSchema):
    id: int


class ServiceLogSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    id: int | None
    service_id: int
    message: str

class ServiceWithLogsSchema(ServiceSchema):
    logs: list[ServiceLogSchema]
