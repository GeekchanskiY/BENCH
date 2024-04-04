import { object, string, number, date } from 'yup';
import { useEffect, useState } from 'react';
import { Formik } from 'formik';

let employeeSchema = object({
  id: number().positive().integer(),
  name: string().required(),
  age: number().required().positive().integer(),
});

async function fetchUsers(){
    let response = await fetch('http://0.0.0.0:3001/v1/employee/')
    let employees = await response.json()
    if (employees == null){
      return []
    }
    let validatedObjects = await Promise.all(employees.map(async (object) => {
        await employeeSchema.validate(object, { abortEarly: false });
        return object;
    }));
    return await validatedObjects
}

function CreateEmployeeForm({refresh, setRefresh}){
    async function createEmployee(data){
        console.log(data)
        try {
            let response = await fetch(
                'http://0.0.0.0:3001/v1/employee/',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        name: data.name,
                        age: data.age
                    })
                }
            )
            //let data = await response.json()
            //console.log(data)
            setRefresh(!refresh)
        } catch (errors) {
            console.log(errors)
            return {
                success: false
            }
        }
        return {
            success: true
        }
    }

    return <Formik
      initialValues={{ name: 'Dmitry', age: 20 }}
      validationSchema={employeeSchema}
      onSubmit={async (values, { setSubmitting, setErrors }) => {

        let data = await createEmployee(values)
        if (data.success == true) {
          setSubmitting(false);
        } else {
          setErrors({ 'email': 'Invalid login or password!' })
        }

      }}>
      {({
        values,
        errors,
        touched,
        handleChange,
        handleBlur,
        handleSubmit,
        isSubmitting,
      }) => (
        <form onSubmit={handleSubmit} className='frm loginform'>
          <h3>Create Employee</h3>
          <input
            type="text"
            name="name"
            onChange={handleChange}
            onBlur={handleBlur}
            value={values.name}
          /> <br />
          <span className='errors'>{errors.name && touched.name && errors.name}</span> <br />
          <input
            type="number"
            name="age"
            onChange={handleChange}
            onBlur={handleBlur}
            value={values.age}
          /> <br />
          <span className='errors'>{errors.age && touched.age && errors.age}</span> <br />
          <button type="submit" disabled={isSubmitting}>
            Create
          </button>
        </form>
      )}
    </Formik>
}


function EmployeeComponent(props){
    async function deleteEmployee(){
      let response = await fetch(
        'http://0.0.0.0:3001/v1/employee/'+props.employee.id,
        {
            method: 'DELETE',
            headers: {
                'Content-Type': 'application/json'
            },
        }
        
      )
      props.setRefresh(!props.refresh)
    }
    return <div className='cv_instance'>
        <ul>
          <li>ID: {props.employee.id}</li>
          <li>Name: {props.employee.name}</li>
          <li>Age: {props.employee.age}</li>
        </ul>
        <button onClick={deleteEmployee}>Delete</button>   
    </div>
}

export default function Employee(){
    const [employees, setEmployees] = useState([])
    const [refresh, setRefresh] = useState(false)
    async function setUp(){
        setEmployees(await fetchUsers())
    }
    useEffect(() => {
        setUp()
    }, [refresh])
    console.log(employees)
    if (employees.length == 0){
      return <div className="cv_model">
        <h1>Employees</h1>
        <span>No records found!</span>
        <CreateEmployeeForm refresh={refresh} setRefresh={setRefresh}/>
      </div>
    } 
    return <div className="cv_model">
        <h1>Employees</h1>
        <div className='cv_instances'>
          {employees.map((employee)=>{
              return <EmployeeComponent employee={employee} key={employee.id} refresh={refresh} setRefresh={setRefresh}/>
          })}
        </div>
        <CreateEmployeeForm refresh={refresh} setRefresh={setRefresh}></CreateEmployeeForm>
    </div>
    
}