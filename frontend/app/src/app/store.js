import { configureStore } from '@reduxjs/toolkit'
import jwtReducer from '../features/jwt/jwtSlice'

export default configureStore({
  reducer: {
    jwt: jwtReducer
  },
})