import { apiSlice } from '../../app/api/apiSlice';

const USER_URL = "/users";

export const draftApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    createBlog: builder.mutation({
      query: (data) => ({
        url: `${USER_URL}/create/blogs`,
        method: 'POST',
        body: {
          ...data
        }
      })
    }),
    updateBlog: builder.mutation({
      query: ({ data, id }) => ({
        url: `${USER_URL}/update/blogs?blogId=${id}`,
        method: 'PUT',
        body: { ...data }
      })
    }),
    deleteBLog: builder.mutation({
      query: (id) => ({
        url: `${USER_URL}/delete/blogs?blogId=${id}`,
        method: 'DELETE',
        credentials: "include",
      })
    }),
    getAllBlogs: builder.query({
      query: () => ({
        url: `${USER_URL}/get/blogs`,
        method: "GET",
        credentials: "include",
      })
    }),
    getBlog: builder.query({
      query: (id) => ({
        url: `${USER_URL}/get/blogs?blogId${id}`,
        method: "GET",
        credentials: "include",
      })
    }),
    getUserBlog: builder.query({
      query: () => ({
        url: `${USER_URL}/get/user/blogs`,
        method: "GET",
        credentials: "include",
      })
    }),
    countMyBlogs: builder.query({
      query: () => ({
        url: `${USER_URL}/count/my/blogs`,
        method: "GET",
        credentials: "include"
      })
    }),
    countAllBlogs: builder.query({
      query: () => ({
        url: `${USER_URL}/count/all/blogs`,
        method: "GET",
        credentials: "include"
      })
    }),
    countUserBlogs: builder.query({
      query: (userId) => ({
        url: `${USER_URL}/count/users/${userId}/blogs`,
        method: "GET",
        credentials: "include"
      })
    }),
  })
})

export const {
  useGetBlogQuery,
  useGetAllBlogsQuery,
  useDeleteBLogMutation,
  useUpdateBlogMutation,
  useCreateBlogMutation,
  useCountUserBlogsQuery,
  useCountMyBlogsQuery,
  useCountAllBlogsQuery
} = draftApiSlice;
