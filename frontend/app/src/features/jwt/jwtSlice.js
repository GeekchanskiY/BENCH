import { createSlice } from '@reduxjs/toolkit'

export const jwtSlice = createSlice({
  name: 'jwt',
  initialState: {
    username: null,
    token: null,
    expires_in: null,
  },
  reducers: {
    logout: (state) => {
      state.username = null
      state.token = null
      state.expires_in = null
    },
    login: (state, action) => {
      state.username = action.payload.username
      state.token = action.payload.token
      state.expires_in = action.payload.expires_in
    },
  },
})

// Action creators are generated for each case reducer function
export const { login, logout } = jwtSlice.actions

export default jwtSlice.reducer