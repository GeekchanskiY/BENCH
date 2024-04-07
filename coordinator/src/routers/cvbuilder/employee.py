from fastapi import APIRouter, Depends, HTTPException, Request, UploadFile, File
from services.cvbuilder.employeeService import EmployeeService
from routers.depends import get_employee_service
from schemas.cvbuilder.employeeSchema import EmployeeSchema, EmployeeDetailedSchema
from schemas.statusSchemas import MessageResponseSchema

from routers.depends import get_jwt_bearer, JWTCredentials
from routers.exceptions import exceptionHandler


router: APIRouter = APIRouter()

@router.get("/{id}", tags=["employees"], response_model=EmployeeDetailedSchema)
async def read_employees(id: int, service = Depends(get_employee_service)):
    try:
        employee = await service.get_employee(id=id)
        return employee
    except Exception as e:
        raise await exceptionHandler(e)
    
@router.post("/create", tags=["employees"], response_model=EmployeeDetailedSchema)
async def create_employee(employee: EmployeeSchema, service: EmployeeService  = Depends(get_employee_service)):
  
    try:
        employee = await service.create_employee(employee=employee)
        return employee
    except Exception as e:
        raise await exceptionHandler(e)
