from pydantic import BaseModel, ConfigDict, validator


class JWTSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    token: str

class JWTDetailedSchema(JWTSchema):
    expires_at: str
    username: str