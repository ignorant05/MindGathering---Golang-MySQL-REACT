import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  blogs: [],
  page: 1,
  size: 10,
  totalBlogs: 0
}

const pageSlice = createSlice({
  name: "page",
  initialState: initialState,
  reducers: {
    setPage: (state, action) => {
      state.page = action.payload.page
    },
    setBlogs: (state, action) => {
      state.blogs = action.payload.blogs
      state.totalBlogs = action.payload.totalBlogs
    }
  }
})


export const { setPage, setBlogs } = pageSlice.actions
export default pageSlice.reducer
