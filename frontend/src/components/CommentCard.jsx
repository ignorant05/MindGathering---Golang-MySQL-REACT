import { useEffect, useRef, useState } from "react";
import { useDeleteCommentMutation, useUpdateCommentMutation, useGetAuthorNameQuery } from "../redux/slices/commentApiSlice";
import { useSelector } from "react-redux";

const CommentCard = (props) => {
  const { cid, author_id, comment } = props.body.comment;
  const blogId = props.body.blogId

  const textareaRef = useRef();
  const [editIt, setEditIt] = useState(false)
  const [mine, setMine] = useState(true)

  const storedID = sessionStorage.getItem('user_id');
  const user = useSelector((state) => state.auth.user);
  const user_id = user?.uid || storedID || "";

  useEffect(() => {
    if (user_id === author_id) {
      setMine(true)
    }
    else {
      setMine(false)
    }
  }, [user_id, author_id])

  const editPressed = () => {
    setEditIt(!editIt)
  }
  const [editCmnt] = useUpdateCommentMutation()
  const [delCmnt] = useDeleteCommentMutation()

  const { data: authorData } = useGetAuthorNameQuery(author_id);
  const username = authorData?.username || "Anonymous";

  const edit = async (e) => {
    e.preventDefault()
    if (editIt) {
      try {
        const updatedContent = textareaRef.current.value;
        await editCmnt({ cid, blogId, content: updatedContent }).unwrap();

      } catch (err) {
        console.error(err)
      }
    }
  }
  const deletePressed = async () => {
    try {
      await delCmnt(cid)
    } catch (err) {
      console.error(err)
    }
  }
  return (
    <>
      <div className="border rounded-md p-3 ml-3 my-3">
        <div className="flex gap-3 items-center">
          <img src="https://avatars.githubusercontent.com/u/22263436?v=4"
            className="object-cover w-8 h-8 rounded-full 
                            border-2 border-emerald-400  shadow-emerald-400
                            "/>
          <h3 className="font-bold">
            {username}
          </h3>
        </div>
        <form onSubmit={edit}>
          <div className="w-full mb-4 border border-gray-200 rounded-lg bg-gray-50 dark:bg-gray-700 dark:border-gray-600">
            <div className="px-4 py-2 bg-white rounded-t-lg dark:bg-gray-800">
              <label htmlFor="comment" className="sr-only">Your comment</label>
              <textarea  {...(!mine || !editIt ? { disabled: true } : {})} defaultValue={comment} ref={textareaRef} id="comment" rows="4" className="w-full px-0 text-sm text-gray-900 bg-white border-0 dark:bg-gray-800 focus:ring-0 dark:text-white dark:placeholder-gray-400" required ></textarea>
            </div>
            <div {...(!mine && { hidden: true })} className="flex items-center justify-between px-3 py-2 border-t dark:border-gray-600 border-gray-200">
              {editIt &&
                <button type="submit" className="inline-flex items-center py-2.5 px-4 text-xs font-medium text-center text-white bg-blue-700 rounded-lg focus:ring-4 focus:ring-blue-200 dark:focus:ring-blue-900 hover:bg-blue-800">
                  Done
                </button>
              }
              <div className="flex ps-0 space-x-1 rtl:space-x-reverse sm:ps-2">
                <button onClick={editPressed} type="button" className="inline-flex justify-center items-center p-2 text-gray-500 rounded-sm cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
                  <span className="sr-only">Edit</span>
                </button>
                <button onClick={deletePressed} type="button" className="inline-flex justify-center items-center p-2 text-gray-500 rounded-sm cursor-pointer hover:text-gray-900 hover:bg-gray-100 dark:text-gray-400 dark:hover:text-white dark:hover:bg-gray-600">
                  <span className="sr-only">Delete</span>
                </button>
              </div>
            </div>
          </div>
        </form >
      </div >
    </>
  )
}

export default CommentCard; 
