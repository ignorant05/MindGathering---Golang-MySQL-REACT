import { apiSlice } from '../../app/api/apiSlice';

const AUTH_URL = "/auth";

export const authApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    register: builder.mutation({
      query: (data) => ({
        url: `${AUTH_URL}/register`,
        method: 'POST',
        body: {
          ...data
        }
      })
    }),
    login: builder.mutation({
      query: (data) => ({
        url: `${AUTH_URL}/login`,
        method: 'POST',
        body: { ...data }
      })
    }),
  })
})

export const {
  useRegisterMutation,
  useLoginMutation,
} = authApiSlice;

