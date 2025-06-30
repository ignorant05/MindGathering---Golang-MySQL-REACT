import { createSlice } from "@reduxjs/toolkit";


const readingBlogSlice = createSlice({
  name: "blogID",
  initialState: {
    blogID: ""
  },
  reducers: {
    setClickedBlog: (state, action) => {
      state.blogID = action.payload.blogID
    },
    clearClickedBlog: (state, action) => {
      state.blogID = ""
    }
  }
})


export const { setClickedBlog, clearClickedBlog } = readingBlogSlice.actions
export default readingBlogSlice.reducer
