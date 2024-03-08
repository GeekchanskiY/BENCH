import { Formik } from 'formik';
import { postRequest } from '../features/requests/requests';
import { useDispatch, useSelector } from "react-redux";
import { login_slice } from '../features/jwt/jwtSlice';
import { useNavigate } from "react-router-dom";
import { useEffect } from 'react';
import * as Yup from 'yup';

const LoginSchema = Yup.object().shape({
    password: Yup.string()
      .min(8, 'Min 8 symbols')
      .max(50, 'Max 50 symbols')
      .required('required field'),
    
    email: Yup.string().email('Wrong e-mail').required('Required field'),
  });


export default function Login(){
    const dispatch = useDispatch()
    const navigate = useNavigate()
    const jwt = useSelector((state) => state.jwt.token)

    useEffect(()=>{
      if (jwt != null){
        navigate('/me')
      }
    }, [])
    
    async function login_request(values){
        console.log(values)
        let data = await postRequest(
            'http://0.0.0.0:80/users/auth',
            {
                "email": values.email,
                "password": values.password
            }
        )
        
        dispatch(
        login_slice({
            username: data.username,
            expires_at: data.expires_at,
            token: data.token
        })
        )
    }
    return <div>
        <Formik
        initialValues={{ email: 'dmt@mail.ru', password: 'password!' }}
        validationSchema={LoginSchema}
        onSubmit={(values, { setSubmitting }) => {
            setTimeout(async () => {
                await login_request(values);
                setSubmitting(false);
                navigate('/me')
              }, 400);
        }}
      >
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
            <h3>Login</h3>
            <input
              type="email"
              name="email"
              onChange={handleChange}
              onBlur={handleBlur}
              value={values.email}
            /> <br />
            <span className='errors'>{errors.email && touched.email && errors.email}</span> <br />
            <input
              type="password"
              name="password"
              onChange={handleChange}
              onBlur={handleBlur}
              value={values.password}
            /> <br />
             <span className='errors'>{errors.password && touched.password && errors.password}</span> <br />
            <button type="submit" disabled={isSubmitting}>
              Submit
            </button>
          </form>
        )}
      </Formik>
    </div>
}