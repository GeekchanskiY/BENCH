import { createSlice } from '@reduxjs/toolkit'

export const jwtSlice = createSlice({
  name: 'jwt',
  initialState: {
    username: "asdasd",
    token: "asdasd",
    expires_in: "asdasdsadasd",
  },
  reducers: {
    logout_slice: (state) => {
      state.username = null
      state.token = null
      state.expires_in = null
    },
    login_slice: (state, action) => {
      state.username = action.payload.username
      state.token = action.payload.token
      state.expires_in = action.payload.expires_in
    },
  },
})

// Action creators are generated for each case reducer function
export const { logout_slice, login_slice } = jwtSlice.actions

export default jwtSlice.reducer