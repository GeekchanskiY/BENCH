import { object, string, number, date } from 'yup';
import { useEffect, useState } from 'react';
import { Formik } from 'formik';

let companySchema = object().shape({
  id: number().positive().integer(),
  name: string().required(),
  rating: number().required(),
  description: string().required(),
  city: string().required(),
  link: string().required(),
});

async function fetchCompanies() {
  let response = await fetch('http://0.0.0.0:3001/v1/company/')
  let companies = await response.json()
  if (companies == null) {
    return []
  }
  let validatedObjects = await Promise.all(companies.map(async (object) => {
    await companySchema.validate(object, { abortEarly: false });
    return object;
  }));
  return await validatedObjects
}


function CreateCompanyForm({ refresh, setRefresh }) {
  async function createCompany(data) {

    try {
      let res = await companySchema.validate(data, { abortEarly: false });
      console.log(res)
      let response = await fetch(
        'http://0.0.0.0:3001/v1/company/',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(res)
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
    initialValues={{
      name: 'Innowise group',
      rating: 5,
      description: 'Best company in the world!',
      city: 'Warsaw',
      link: 'https://innowise.com/'
    }}
    validationSchema={companySchema}
    onSubmit={async (values, { setSubmitting, setErrors }) => {

      let data = await createCompany(values)
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
        <h3>Create Company</h3>
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
          name="rating"
          onChange={handleChange}
          onBlur={handleBlur}
          value={values.rating}
        /> <br />
        <span className='errors'>{errors.rating && touched.rating && errors.rating}</span> <br />
        <input type="text" name="description" onChange={handleChange} onBlur={handleBlur} value={values.description} /> <br />
        <span className='errors'>{errors.description && touched.description && errors.description}</span> <br />
        <input type="text" name="city" onChange={handleChange} onBlur={handleBlur} value={values.city} /> <br />
        <span className='errors'>{errors.city && touched.city && errors.city}</span> <br />
        <input type="text" name="link" onChange={handleChange} onBlur={handleBlur} value={values.link} /> <br />
        <span className='errors'>{errors.link && touched.link && errors.link}</span> <br />
        <button type="submit" disabled={isSubmitting}>
          Create
        </button>
      </form>
    )}
  </Formik>
}

function CompanyComponent(props) {
  async function deleteCompany() {
    let response = await fetch(
      'http://0.0.0.0:3001/v1/company/' + props.company.id,
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
      <li>ID: {props.company.id}</li>
      <li>Name: {props.company.name}</li>
      <li>Rating: {props.company.rating}</li>
      <li>Description: {props.company.description}</li>
      <li>City: {props.company.city}</li>
      <li>Link: {props.company.link}</li>
    </ul>
    <button onClick={deleteCompany}>Delete</button>
  </div>
}


export default function Company() {
  const [companies, setCompanies] = useState([])
  const [refresh, setRefresh] = useState(false)
  async function setUp() {
    setCompanies(await fetchCompanies())
  }
  useEffect(() => {
    setUp()
  }, [refresh])
  console.log(companies)
  if (companies.length == 0) {
    return <div className="cv_model">
      <h1>Companies</h1>
      <span>No records found!</span>
      <CreateCompanyForm refresh={refresh} setRefresh={setRefresh} />
    </div>
  }
  return <div className="cv_model">
    <h1>Companies</h1>
    <div className='cv_instances'>
      {companies.map((domain) => {
        return <CompanyComponent company={domain} key={"company_" + domain.id} refresh={refresh} setRefresh={setRefresh} />
      })}
    </div>
    <CreateCompanyForm refresh={refresh} setRefresh={setRefresh}></CreateCompanyForm>
  </div>
}