import { createSlice } from "@reduxjs/toolkit";

const initialState = {
  draft: {
    title: "",
    content: ""
  }
}

const draftSlice = createSlice({
  name: "draft",
  initialState: initialState,
  reducers: {
    setDraft: (state, action) => {
      state.draft = { ...state.draft, ...action.payload }
    },
    clearDraft: (state, action) => {
      state.draft = { title: "", content: "" };
    }
  }
})


export const { setDraft, clearDraft } = draftSlice.actions
export default draftSlice.reducer
