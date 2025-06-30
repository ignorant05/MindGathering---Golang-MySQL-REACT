import { createSlice } from "@reduxjs/toolkit";

const init = {
  user: null,
  access_token: null,
}

const authSlice = createSlice({
  name: "auth",
  initialState: init,
  reducers: {
    setInfo: (state, action) => {
      const { user, access_token } = action.payload
      state.user = user
      state.access_token = access_token
    },
    logOut: (state, action) => {
      state.user = null
      state.access_token = null
    }
  }
})


export const { setInfo, logOut } = authSlice.actions
export default authSlice.reducer
