from pydantic import BaseModel, ConfigDict

class MessageResponseSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    msg: str