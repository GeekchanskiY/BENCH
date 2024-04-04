import { object, string, number, date } from 'yup';
import { useEffect, useState } from 'react';
import { Formik } from 'formik';

let domainSchema = object({
  id: number().positive().integer(),
  name: string().required(),
});

async function fetchDomain(){
    let response = await fetch('http://0.0.0.0:3001/v1/domain/')
    let domains = await response.json()
    if (domains == null){
      return []
    }
    let validatedObjects = await Promise.all(domains.map(async (object) => {
        await domainSchema.validate(object, { abortEarly: false });
        return object;
    }));
    return await validatedObjects
}

function CreateDomainForm({refresh, setRefresh}){
    async function createDomain(data){
        console.log(data)
        try {
            let response = await fetch(
                'http://0.0.0.0:3001/v1/domain/',
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        name: data.name
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
      initialValues={{ name: 'DevOps' }}
      validationSchema={domainSchema}
      onSubmit={async (values, { setSubmitting, setErrors }) => {

        let data = await createDomain(values)
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
          <h3>Create Domain</h3>
          <input
            type="text"
            name="name"
            onChange={handleChange}
            onBlur={handleBlur}
            value={values.name}
          /> <br />
          <span className='errors'>{errors.name && touched.name && errors.name}</span> <br />
          <button type="submit" disabled={isSubmitting}>
            Create
          </button>
        </form>
      )}
    </Formik>
}


function DomainComponent(props){
    async function deleteDomain(){
      let response = await fetch(
        'http://0.0.0.0:3001/v1/domain/'+props.domain.id,
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
          <li>ID: {props.domain.id}</li>
          <li>Name: {props.domain.name}</li>
        </ul>
        <button onClick={deleteDomain}>Delete</button>   
    </div>
}

export default function Domain(){
    const [domains, setDomains] = useState([])
    const [refresh, setRefresh] = useState(false)
    async function setUp(){
        setDomains(await fetchDomain())
    }
    useEffect(() => {
        setUp()
    }, [refresh])
    console.log(domains)
    if (domains.length == 0){
      return <div className="cv_model">
        <h1>Domains</h1>
        <span>No records found!</span>
        <CreateDomainForm refresh={refresh} setRefresh={setRefresh}/>
      </div>
    } 
    return <div className="cv_model">
        <h1>Domains</h1>
        <div className='cv_instances'>
          {domains.map((domain)=>{
              return <DomainComponent domain={domain} key={"domain_"+domain.id} refresh={refresh} setRefresh={setRefresh}/>
          })}
        </div>
        <CreateDomainForm refresh={refresh} setRefresh={setRefresh}></CreateDomainForm>
    </div>
    
}