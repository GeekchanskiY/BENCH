import { Formik } from 'formik';
import { useSelector } from "react-redux";
import * as Yup from 'yup';
import { postRequestAuth } from '../../features/requests/requests';

const ServiceSchema = Yup.object().shape({
    name: Yup.string().required(),

    
    description: Yup.string().required(),
    is_active: Yup.boolean().required(),
    url: Yup.string().url().required(),
    ping_url: Yup.string().url().required()
  });


export default function CreateService(props){
    const jwt = useSelector((state) => state.jwt.token)

    async function createServiceRequest(values){
        let data = await postRequestAuth(
            'http://0.0.0.0:80/services/create',
            {
                "name": values.name,
                "description": values.description,
                "is_active": values.is_active,
                "url": values.url,
                "ping_url": values.ping_url
            },
            jwt
        )
        if (data.success){
            props.setReload(prev => !prev)
        } else {
            alert('Error in request')
        }
    }
    return <Formik
        initialValues={{
            name: 'Coordinator',
            description: 'Api Gateway',
            is_active: true,
            url: "http://0.0.0.0:80/",
            ping_url: "http://0.0.0.0:80/ping",
        }}
        validationSchema={ServiceSchema}
        onSubmit={async (values, { setSubmitting, setErrors }) => {
          
          let data = await createServiceRequest(values);
          if (data.success == true){
            setSubmitting(false);
          } else {
            setErrors({'name': 'Error occured!'})
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
            <h3>Add service</h3>
            <input
              type="text"
              name="name"
              onChange={handleChange}
              onBlur={handleBlur}
              value={values.name}
            /> <br />
            <span className='errors'>{errors.name && touched.name && errors.name}</span> <br />
            <input
              type="text"
              name="description"
              onChange={handleChange}
              onBlur={handleBlur}
              value={values.description}
            /> <br />
             <span className='errors'>{errors.description && touched.description && errors.description}</span> <br />
            <input 
                type="checkbox"
                name="is_active"
                onChange={handleChange}
                onBlur={handleBlur}
                checked={values.is_active}
            /> <br />
            <input
              type="text"
              name="url"
              onChange={handleChange}
              onBlur={handleBlur}
              value={values.url}
            /> <br />
             <span className='errors'>{errors.url && touched.url && errors.url}</span> <br />
             <input
              type="text"
              name="url"
              onChange={handleChange}
              onBlur={handleBlur}
              value={values.ping_url}
            /> <br />
             <span className='errors'>{errors.ping_url && touched.ping_url && errors.ping_url}</span> <br />

            <button type="submit" disabled={isSubmitting}>
              Submit
            </button>
          </form>
        )}
    </Formik>
}