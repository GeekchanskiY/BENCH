from schemas.employeeSchema import EmployeeFullSchema, EmployeeSchema
import aiohttp

class EmployeeService:
    def __init__(self) -> None:
        self.session = aiohttp.ClientSession()

    async def get_employee(self, id):
        response = await self.session.get(f"http://finance:3001/v1/employee/get/{id}")
        data: dict = await response.json()
        data = data['Data']
        return EmployeeFullSchema(**data)
    
    async def create_employee(self, employee: EmployeeSchema):
        response = await self.session.post(f"http://finance:3001/v1/employee/create", json=dict(employee))
        data: dict = await response.json()
        print(data)
        data = data['Data']
        return EmployeeFullSchema(**data)