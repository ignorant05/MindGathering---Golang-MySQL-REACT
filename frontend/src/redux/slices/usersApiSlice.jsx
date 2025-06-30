import { apiSlice } from "../../app/api/apiSlice";

const USER_URL = "/users"

export const usersApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    logout: builder.mutation({
      query: () => ({
        url: `${USER_URL}/logout`,
        method: 'POST',
        credentials: "include",
      })
    }),
    delete: builder.mutation({
      query: () => ({
        url: `${USER_URL}/delete`,
        method: 'DELETE',
        credentials: "include",
      })
    }),
    update: builder.mutation({
      query: (data) => ({
        url: `${USER_URL}/update/credentials`,
        method: "PUT",
        credentials: "include",
        body: {
          ...data
        }
      })
    }),
    updatePic: builder.mutation({
      query: (data) => ({
        url: `${USER_URL}/update/picture`,
        method: "PUT",
        credentials: "include",
        body: {
          ...data
        }
      })
    }),
    getMe: builder.query({
      query: () => ({
        url: `${USER_URL}/get/profile`,
        method: "GET",
        credentials: "include"
      })
    }),
    getUserInfo: builder.query({
      query: (userId) => ({
        url: `${USER_URL}/profile?userId=${userId}`,
        method: "GET",
        credentials: "include"
      })
    }),
    getMyPic: builder.query({
      query: () => ({
        url: `${USER_URL}/get/my/image`,
        method: "GET",
        credentials: "include"
      })
    }),
    getUserPic: builder.query({
      query: (userId) => ({
        url: `${USER_URL}/get/users/${userId}/image`,
        method: "GET",
        credentials: "include"
      })
    }),
    pagination: builder.query({
      query: ({ page, size }) => ({
        url: `${USER_URL}/pages?page=${page}&size=${size}`,
        method: "GET",
        credentials: "include",
      })
    }),
    countAllBlogs: builder.query({
      query: () => ({
        url: `${USER_URL}/count/all/blogs`,
        method: "GET",
        credentials: "include",
      })
    }),
  })
})

export const {
  useGetMeQuery,
  useLogoutMutation,
  useDeleteMutation,
  useUpdateMutation,
  useUpdatePicMutation,
  useGetUserInfoQuery,
  useGetMyPicQuery,
  useGetUserPicQuery,
  usePaginationQuery,
  useCountAllBlogsQuery
} = usersApiSlice;
