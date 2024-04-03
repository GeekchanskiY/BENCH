import { object, string, number, date } from 'yup';
import { useEffect, useState } from 'react';

let employeeSchema = object({
  id: number().positive().integer(),
  name: string().required(),
  age: number().required().positive().integer(),
});

async function fetchUsers(){
    let response = await fetch('http://0.0.0.0:3001/v1/employee/')
    let employees = await response.json()
    let validatedObjects = await Promise.all(employees.map(async (object) => {
        await employeeSchema.validate(object, { abortEarly: false });
        return object;
    }));
    return await validatedObjects
}

export default function Employee(){
    const [employees, setEmployees] = useState([
        {name: 'Dima', age: 12}
    ])

    useEffect(async () => {
        setEmployees(await fetchUsers())
    }, [])
    
    return <div className="cv_model">
        <h1>Employees</h1>
        {employees.map((employee)=>{
            return <span>{employee.name}</span>
        })}
    </div>
}