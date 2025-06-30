import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  comments: []
}

const commentSlice = createSlice({
  name: 'comments',
  initialState: initialState,
  reducers: {
    setComments: (state, action) => {
      state.comments = { ...action.payload.comments }
    },
    clearComments: (state, action) => {
      state.comments = []
    }
  }
})


export const { setComments, clearComments } = commentSlice.actions
export default commentSlice.reducer
