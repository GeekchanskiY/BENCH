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
    let validatedObjects = await Promise.all(employees.map(async (object) => {
        await employeeSchema.validate(object, { abortEarly: false });
        return object;
    }));
    return await validatedObjects
}

function CreateEmployeeForm(){
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
            // let data = await response.json()
            // console.log(data)
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
      initialValues={{ name: 'Dmitry', age: '12' }}
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
    return <div className='cv_instance'>
        <span>ID: {props.employee.id}</span> <br />
        <span>Name: {props.employee.name}</span> <br />
        <span>Age: {props.employee.age}</span>    
    </div>
}

export default function Employee(){
    const [employees, setEmployees] = useState([])
    async function setUp(){
        setEmployees(await fetchUsers())
    }
    useEffect(() => {
        setUp()
    }, [])
    
    return <div className="cv_model">
        <h1>Employees</h1>
        {employees.map((employee)=>{
            return <EmployeeComponent employee={employee}/>
        })}
        <CreateEmployeeForm></CreateEmployeeForm>
    </div>
}