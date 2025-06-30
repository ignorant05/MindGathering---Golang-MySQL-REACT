import { useEffect, useRef } from "react"
import CommentCard from "../components/CommentCard"
import { useCommentOnBLogMutation, useGetBlogCommentsQuery } from "../redux/slices/commentApiSlice"
import { useSelector } from "react-redux"
import { setInfo } from "../redux/slices/authSlice"
import { useGetBlogQuery } from "../redux/slices/draftApiSlice"

const BlogViewScreen = () => {
  const storedToken = sessionStorage.getItem('access_token');
  const auth = useSelector((state) => state.auth);
  const access_token = auth?.access_token || storedToken || "";
  const blogId = useSelector((state) => state.blogID.blogID) || "";
  if (blogId === "") {
    console.log("error")
  }

  const { data: blogData } = useGetBlogQuery(blogId);
  const { data: comments } = useGetBlogCommentsQuery(blogId);

  useEffect(() => {
    setInfo({ access_token })
  }, [access_token])

  const commentRef = useRef("")
  const [commentOnBlog] = useCommentOnBLogMutation()

  const createNewComment = async (e) => {
    e.preventDefault()
    const comment = commentRef?.current?.value
    try {
      await commentOnBlog(comment).unwrap()
    } catch (err) {
      console.log(err)
    }
  }

  return (
    <>
      <div className="h-screen flex justify-center flex-col">
        <div className="w-1/2 ">
          <span className="py-10 mt-10 flex justify-center flex-col font-mono text-7xl">
            {blogData.title}
          </span>

        </div>
        <div className="py-8 px-8 ">
          <p className="py-8  text-3xl">
            {blogData.content}
          </p>
        </div>
        <div className="flex items-center py-8 px-8">
          <p className="py-8"><span className="font-mono font-bold text-3xl">Created By  </span>
            <span className="font-mono font-thin text-1xl">author link</span>
          </p>
        </div>
        <div className="w-full bg-white rounded-lg border p-2 my-4 mx-6">

          <h3 className="font-bold">Discussion</h3>

          <form onSubmit={createNewComment}>
            <div className="w-full px-3 my-2">
              <textarea
                ref={commentRef}
                className="bg-gray-100 rounded border border-gray-400 leading-normal resize-none w-full h-20 py-2 px-3 font-medium placeholder-gray-700 focus:outline-none focus:bg-white"
                name="body" placeholder='Type Your Comment' required></textarea>
            </div>

            <div className="w-full flex justify-end px-3">
              <input type='submit' className="px-2.5 py-1.5 rounded-md text-white text-sm bg-indigo-500" value='Post Comment' />
            </div>

            <div className="flex flex-col">
              <div className="border rounded-md p-3 ml-3 my-3">
                <ul>
                  <li>
                    {comments.map((comment) => <CommentCard key={comment.id} body={{ comment, blogId }} />)}
                  </li>
                </ul>
              </div>
            </div>
          </form>
        </div>
      </div>
    </>
  )
}

export default BlogViewScreen
