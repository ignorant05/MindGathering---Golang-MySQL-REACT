import { apiSlice } from '../../app/api/apiSlice';

const USER_URL = "/users";

export const commentApiSlice = apiSlice.injectEndpoints({
  endpoints: (builder) => ({
    commentOnBLog: builder.mutation({
      query: (data) => ({
        url: `${USER_URL}/comment`,
        method: 'POST',
        body: {
          ...data
        }
      })
    }),
    updateComment: builder.mutation({
      query: (data) => ({
        url: `${USER_URL}/update/blogs/${data.blogId}/comments?commentId=${data.commentId}`,
        method: 'PUT',
        body: { ...data.content }
      })
    }),
    deleteComment: builder.mutation({
      query: (commentId) => ({
        url: `${USER_URL}/delete/comments?commentId=${commentId}`,
        method: 'DELETE',
        credentials: "include",
      })
    }),
    getAllComments: builder.query({
      query: () => ({
        url: `${USER_URL}/get/comments`,
        method: "GET",
        credentials: "include",
      })
    }),
    getMyComments: builder.query({
      query: () => ({
        url: `${USER_URL}/get/user/comments`,
        method: "GET",
        credentials: "include",
      })
    }),
    getBlogComments: builder.query({
      query: (blogId) => ({
        url: `${USER_URL}/get/blogs/${blogId}/comments`,
        method: "GET",
        credentials: "include",
      })
    }),
    getAuthorName: builder.query({
      query: (userId) => ({
        url: `${USER_URL}/get/comments/users?userId=${userId}`,
        method: "GET",
        credentials: "include",
      })
    }),
    countMyComments: builder.query({
      query: () => ({
        url: `${USER_URL}/count/my/comments`,
        method: "GET",
        credentials: "include"
      })
    }),
    countAllComments: builder.query({
      query: () => ({
        url: `${USER_URL}/count/all/comments`,
        method: "GET",
        credentials: "include"
      })
    }),
    countUserComments: builder.query({
      query: (userId) => ({
        url: `${USER_URL}/count/users/${userId}/comments`,
        method: "GET",
        credentials: "include"
      })
    }),
  })
})

export const {
  useGetMyCommentsQuery,
  useGetAllCommentsQuery,
  useCommentOnBLogMutation,
  useDeleteCommentMutation,
  useUpdateCommentMutation,
  useGetBlogCommentsQuery,
  useGetAuthorNameQuery,
  useCountMyCommentsQuery,
  useCountUserCommentsQuery,
  useCountAllCommentsQuery
} = commentApiSlice;
