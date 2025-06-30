import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { setInfo, logOut } from '../../redux/slices/authSlice'

const AUTH_URL = "/api/v1";

const baseQuery = fetchBaseQuery({
  baseUrl: AUTH_URL,
  credentials: "include",
  prepareHeaders: (headers, { getState }) => {
    const access_token = getState().auth.access_token
    if (access_token) {
      headers.set("authorization", `Bearer ${access_token}`);
    }
    return headers
  }
})


const reAuthBaseQuery = async (api, args, extraOptions) => {
  let result = await baseQuery(api, args, extraOptions)

  if (result?.error?.status === 401) {

    console.log("Sending refresh token...");

    const refreshResult = baseQuery({
      url: "/auth/refresh",
      method: "POST"
    }, args, extraOptions)
    console.log("refresh result: ", refreshResult);
    if (refreshResult?.data) {
      const user = api.getState().user
      api.dispatch(setInfo({ ...refreshResult.data, user }))
      result = await baseQuery(api, args, extraOptions)
    } else {
      api.dispatch(logOut())
    }

    return result

  }
}

export const apiSlice = createApi({
  baseQuery: reAuthBaseQuery,
  endpoints: (builder) => ({})
})  
