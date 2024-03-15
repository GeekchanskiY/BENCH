import { Formik } from 'formik';
import { useSelector } from "react-redux";
import * as Yup from 'yup';
import { postRequestAuth } from '../../features/requests/requests';
import '../../styles/services.css'

const ServiceSchema = Yup.object().shape({
  name: Yup.string().required(),


  description: Yup.string().required(),
  is_active: Yup.boolean().required(),
  url: Yup.string().url().required(),
  ping_url: Yup.string().url().required(),

  image: Yup.mixed()
});


export default function CreateService(props) {
  const jwt = useSelector((state) => state.jwt.token)

  async function createServiceRequest(values) {
    let data = await postRequestAuth(
      'http://0.0.0.0:80/services/create',
      {
        "name": values.name,
        "description": values.description,
        "is_active": values.is_active,
        "url": values.url,
        "ping_url": values.ping_url,
      },
      jwt
    )

    if (data.success) {

      if (values.image != undefined) {
        let formData = new FormData();
        let file = document.getElementById('imageFileInput').files[0]
        formData.append('image', file)

        fetch(
          'http://0.0.0.0:80/services/service/' + data.response.id + '/upload_image',
          {
            method: 'POST',
            headers: {
              'Authorization': 'Bearer ' + jwt
            },
            body: formData
          }
        )
          .then(res => {
            if (res.status != 200) {
              alert('Image not uploaded!')
            }
            console.log(res.status)

          })
      }
      props.setReload(prev => !prev);
      return data;

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
      if (data.success == true) {
        setSubmitting(false);
      } else {
        setErrors({ 'name': 'Error occured!' })
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
      <form onSubmit={handleSubmit} className='form form-create-service'>
        <h3>Add service</h3>
        <input
          type="text"
          name="name"
          onChange={handleChange}
          onBlur={handleBlur}
          value={values.name}
        />
        <span className='errors'>{errors.name && touched.name}</span> 
        <input
          type="text"
          name="description"
          onChange={handleChange}
          onBlur={handleBlur}
          value={values.description}
        /> 
        <span className='errors'>{errors.description && touched.description}</span> 
        
        <label className='checkbox'>
          Is Active?
          <input
            type="checkbox"
            name="is_active"
            onChange={handleChange}
            onBlur={handleBlur}
            checked={values.is_active}
          /> 
        </label>
        
        <input
          type="text"
          name="url"
          onChange={handleChange}
          onBlur={handleBlur}
          value={values.url}
        /> 
        <span className='errors'>{errors.url && touched.url}</span> 
        <input
          type="text"
          name="url"
          onChange={handleChange}
          onBlur={handleBlur}
          value={values.ping_url}
        /> 
        <span className='errors'>{errors.ping_url && touched.ping_url}</span>
        <label class="file-upload">
          <input
            type="file"
            name="image"
            id='imageFileInput'
            accept="image/png, image/jpeg"
            onChange={handleChange}
            onBlur={handleBlur}
            value={values.image}
          /> 
          Service image
        </label>
        
        <span className='errors'>{errors.image && touched.image}</span> 
        <input type="submit" disabled={isSubmitting} content='Submit'/>
          
      </form>
    )}
  </Formik>
}