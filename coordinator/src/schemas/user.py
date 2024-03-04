from pydantic import BaseModel, ConfigDict

class User(BaseModel):
    model_config = ConfigDict(from_attributes=True)

    id: int
    name: str
    email: str
    
    is_staff: bool
    ip_adress: str

class UserPrivate(BaseModel):
    password: str