import { useEffect, useState } from "react"
import { useDispatch, useSelector } from "react-redux"
import { clearDraft, setDraft } from "../redux/slices/draftSlice";
import { setInfo } from "../redux/slices/authSlice";
import { useUpdateBlogMutation, useGetBlogQuery } from "../redux/slices/draftApiSlice";
import { useLocation } from "react-router-dom";
import { setClickedBlog } from "../redux/slices/readBlogSlice";

const EditBlogScreen = () => {
  const [title, setTitle] = useState("")
  const [content, setContent] = useState("")

  const location = useLocation();
  const params = new URLSearchParams(location.search);
  const blogId = params.get("blogId");


  const { data: blogData } = useGetBlogQuery(blogId);

  const storedToken = sessionStorage.getItem('access_token');
  const auth = useSelector((state) => state.auth);
  const access_token = auth?.access_token || storedToken || "";


  useEffect(() => {
    if (blogData) {
      setTitle(blogData.title || "");
      setContent(blogData.content || "");
      dispatch(setClickedBlog(blogId));
    }
  }, [blogData, blogId, dispatch]);


  useEffect(() => {
    setInfo({ access_token })
  }, [access_token])


  const dispatch = useDispatch()

  const [edit] = useUpdateBlogMutation()

  const editBlog = async (e) => {
    e.preventDefault()

    const payload = {
      title, content
    }

    try {
      await edit({ data: payload, id: blogId }).unwrap()
      dispatch(setDraft({ title, content }));
    } catch (err) {
      console.log(err)
      dispatch(clearDraft())
    }
  }


  return (
    <>
      <div className="w-full bg-white rounded-lg border p-2 my-4 h-screen">
        <form className="h-full" onSubmit={editBlog}>
          <h3 className="font-bold">Edit Blog</h3>
          <div className="w-full px-3 my-2">
            <textarea
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="bg-gray-100 rounded border border-gray-400 leading-normal resize-none w-full h-20 py-2 px-3 font-medium placeholder-gray-700 focus:outline-none focus:bg-white"
              name="title" placeholder='You blog title here ' required></textarea>
          </div>
          <div className="h-full">
            <div className="px-3 my-2 h-full">
              <textarea
                value={content}
                onChange={(e) => setContent(e.target.value)}
                className="bg-gray-100 rounded border border-gray-400 leading-normal resize-none w-full h-full py-2 px-3 font-medium placeholder-gray-700 focus:outline-none focus:bg-white"
                name="content" placeholder='Blog content ...' required></textarea>
            </div>

            <div className="w-full flex justify-end px-3">
              <input type='submit' className="px-2.5 py-1.5 rounded-md text-white text-sm bg-indigo-500" value='Create Blog' />
            </div>
            <div className="px-3 my-2 h-full pb-44">
            </div>
          </div>
        </form>
      </div>
    </>
  )
}

export default EditBlogScreen
