from pydantic import BaseModel, ConfigDict, validator

class UserSchema(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    id: int | None = None
    name: str
    email: str
    
    is_staff: bool
    ip_adress: str

class RegisterUser(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)
    name: str = "SampleUserName"
    email: str = "SampleUserEmail"
    password: str = "SampleStrongPassword"

    @validator('password')
    def validate_password(cls, v: str, values: dict, **kwargs):
        if len(v) <= 8:
            raise ValueError('Password is too short!')
        
        return v

class LoginUser(BaseModel):
    model_config: ConfigDict = ConfigDict(from_attributes=True)

    email: str = "SampleUserEmail"
    password: str = "SampleStrongPassword"

    @validator('password')
    def validate_password(cls, v: str, values: dict, **kwargs):
        if len(v) <= 8:
            raise ValueError('Password is too short!')
        
        return v


class UserPrivateSchema(UserSchema):
    password: str