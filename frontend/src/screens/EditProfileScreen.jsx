import { useEffect, useRef, useState } from "react"
import { useDeleteMutation, useGetMyPicQuery, useLogoutMutation, useUpdateMutation, useUpdatePicMutation } from "../redux/slices/usersApiSlice"
import { useNavigate } from "react-router-dom"
import { useDispatch, useSelector } from "react-redux"
import { setInfo } from "../redux/slices/authSlice"
import fallbackImage from "../assets/default_picture.png"

const EditProfileScreen = () => {
  const usernameRef = useRef("")
  const emailRef = useRef("")
  const currentPasswordRef = useRef("")
  const newPasswordRef = useRef("")
  const confirmNewPasswordRef = useRef("")
  const [picture, setPicture] = useState("../assets/default_picture.png")

  const { data: profilePic } = useGetMyPicQuery();

  const base64ImageSrc = profilePic?.data
    ? `data:${profilePic.type};base64,${profilePic.data}`
    : fallbackImage;
  if (base64ImageSrc) {
    setPicture(base64ImageSrc)
  }

  const navigate = useNavigate()
  const dispatch = useDispatch();

  const storedToken = sessionStorage.getItem('access_token');
  const auth = useSelector((state) => state.auth);
  const access_token = auth?.access_token || storedToken || "";
  const user = auth?.user || {};

  useEffect(() => {
    setInfo({ access_token })
  }, [access_token])

  const [update] = useUpdateMutation()
  const [fullUpdate] = useUpdatePicMutation()
  const [del] = useDeleteMutation()
  const [logout] = useLogoutMutation()

  var loadFile = function(event) {
    const file = event.target.files[0];

    if (file) {
      const imageUrl = URL.createObjectURL(file);
      setPicture(imageUrl);
      var output = document.getElementById('preview_img');
      output.src = URL.createObjectURL(file)
      output.onload = function() {
        URL.revokeObjectURL(output.src)
      }
    }
  };
  const deleteAccount = async () => {
    try {
      await del().unwrap()
      navigate("/auth/register")
    } catch (err) {
      console.log("Internale server error")
      console.error(err?.data?.message || err?.error)
    }
  }

  const logOut = async () => {
    try {
      await logout().unwrap()
      navigate("/auth/login")
    } catch (err) {
      console.log("Internale server error")
      console.error(err?.data?.message || err?.error)
    }
  }

  const updateOnSubmit = async (e) => {
    e.preventDefault()

    const newUsername = usernameRef?.current?.value
    const newEmail = emailRef?.current?.value
    const currPass = currentPasswordRef?.current?.value
    const newPass = newPasswordRef?.current?.value
    const confirmNewPass = confirmNewPasswordRef?.current?.value

    if (newPass !== confirmNewPass) {
      console.error("Password don't match")
      return
    }

    const payload = {
      newUsername,
      newEmail,
      currPass,
      newPass,
      picture
    }

    try {
      let updatedUser

      if (payload.picture) {
        updatedUser = await fullUpdate(payload).unwrap()
      } else {
        updatedUser = await update(payload).unwrap()
      }
      sessionStorage.setItem("user_id", updatedUser.uid)

      if (updatedUser && access_token) {
        navigate("/auth/profile")
        dispatch(setInfo({ user: updatedUser, access_token: access_token }));
      }

    } catch (err) {
      console.log("Internale server error")
      console.error(err?.data?.message || err?.error)
    }
  }

  return (
    <>
      <div className="flex items-center lg:w-full space-x-2 ">
        <div className="flex justify-center items-center text-black h-11/12 lg:w-full space-y-24 box-border px-4 mx-auto border-solid sm:px-6 md:px-6 lg:px-0 max-w-7xl">
          <div className="flex items-centes justify-center my-20 px-6 text-center  h-full rounded-lg lg:mt-0 xl:px-10">
            <div className="w-96 px-6 py-6 text-center rounded-lg lg:mt-0 xl:px-10">
              <div className="space-y-4 xl:space-y-6">
                <form >
                  <div className="flex items-center space-x-6">
                    <div className="shrink-0">
                      <img id='preview_img' className="h-16 w-16 object-cover rounded-full" src={picture} alt="Current profile photo" />
                    </div>
                    <label className="block">
                      <span className="sr-only">Choose profile photo</span>
                      <input type="file" onChange={loadFile} className="block w-full text-sm text-slate-500
                            file:mr-4 file:py-2 file:px-4
                            file:rounded-full file:border-0
                            file:text-sm file:font-semibold
                            file:bg-violet-50 file:text-violet-700
                            hover:file:bg-violet-100
                          "/>
                    </label>
                  </div>
                </form>
                <div className="space-y-2">
                  <div className="flex justify-center items-center flex-col space-y-3 text-lg font-medium leading-6">
                    <h3 className="text-white">{user.username}</h3>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="w-full bg-gray-100 lg:w-1/2 lg:h-full flex items-center justify-center">
            <div className="max-w-md w-full p-6">
              <h1 className="text-3xl font-semibold mb-6 text-black text-center">Profile</h1>
              <div className="mt-4 flex flex-col lg:flex-row items-center justify-between">
                <div className="w-full lg:w-1/2 mb-2 lg:mb-0">
                  <button type="button" className="w-full flex justify-center items-center gap-2 bg-white text-sm text-gray-600 p-2 rounded-md hover:bg-gray-50 border border-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-200 transition-colors duration-300">
                    <svg rect="http://www.w3.org/2000/svg" viewBox="0 0 512 512" className="w-4" id="google">
                      <path fill="#fbbb00" d="M113.47 309.408 95.648 375.94l-65.139 1.378C11.042 341.211 0 299.9 0 256c0-42.451 10.324-82.483 28.624-117.732h.014L86.63 148.9l25.404 57.644c-5.317 15.501-8.215 32.141-8.215 49.456.002 18.792 3.406 36.797 9.651 53.408z"></path>
                      <path fill="#518ef8" d="M507.527 208.176C510.467 223.662 512 239.655 512 256c0 18.328-1.927 36.206-5.598 53.451-12.462 58.683-45.025 109.925-90.134 146.187l-.014-.014-73.044-3.727-10.338-64.535c29.932-17.554 53.324-45.025 65.646-77.911h-136.89V208.176h245.899z"></path>
                      <path fill="#28b446" d="m416.253 455.624.014.014C372.396 490.901 316.666 512 256 512c-97.491 0-182.252-54.491-225.491-134.681l82.961-67.91c21.619 57.698 77.278 98.771 142.53 98.771 28.047 0 54.323-7.582 76.87-20.818l83.383 68.262z"></path>
                      <path fill="#f14336" d="m419.404 58.936-82.933 67.896C313.136 112.246 285.552 103.82 256 103.82c-66.729 0-123.429 42.957-143.965 102.724l-83.397-68.276h-.014C71.23 56.123 157.06 0 256 0c62.115 0 119.068 22.126 163.404 58.936z"></path>
                    </svg> Sign Up with Google </button>
                </div>
                <div className="w-full lg:w-1/2 ml-0 lg:ml-2">
                  <button type="button" className="w-full flex justify-center items-center gap-2 bg-white text-sm text-gray-600 p-2 rounded-md hover:bg-gray-50 border border-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-200 transition-colors duration-300">
                    <svg rect="http://www.w3.org/2000/svg" viewBox="0 0 16 16" id="github" className="w-4">
                      <path d="M7.999 0C3.582 0 0 3.596 0 8.032a8.031 8.031 0 0 0 5.472 7.621c.4.074.546-.174.546-.387 0-.191-.007-.696-.011-1.366-2.225.485-2.695-1.077-2.695-1.077-.363-.928-.888-1.175-.888-1.175-.727-.498.054-.488.054-.488.803.057 1.225.828 1.225.828.714 1.227 1.873.873 2.329.667.072-.519.279-.873.508-1.074-1.776-.203-3.644-.892-3.644-3.969 0-.877.312-1.594.824-2.156-.083-.203-.357-1.02.078-2.125 0 0 .672-.216 2.2.823a7.633 7.633 0 0 1 2.003-.27 7.65 7.65 0 0 1 2.003.271c1.527-1.039 2.198-.823 2.198-.823.436 1.106.162 1.922.08 2.125.513.562.822 1.279.822 2.156 0 3.085-1.87 3.764-3.652 3.963.287.248.543.738.543 1.487 0 1.074-.01 1.94-.01 2.203 0 .215.144.465.55.386A8.032 8.032 0 0 0 16 8.032C16 3.596 12.418 0 7.999 0z"></path>
                    </svg> Sign Up with Github </button>
                </div>
              </div>
              <div className="mt-4 text-sm text-gray-600 text-center">
                <p>or with email</p>
              </div>
              <form onSubmit={updateOnSubmit} action="#" method="POST" className="space-y-4">
                <div>
                  <label htmlFor="username" className="block text-sm font-medium text-gray-700">Username</label>
                  <input defaultValue={user.useraneme || "Not sat"}
                    ref={usernameRef} type="text" id="username" name="username" className="mt-1 p-2 w-full border rounded-md focus:border-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-300 transition-colors duration-300" />
                </div>
                <div>
                  <label htmlFor="email" className="block text-sm font-medium text-gray-700">Email</label>
                  <input defaultValue={user.email || "Not sat"}
                    ref={emailRef} type="text" id="email" name="email" className="mt-1 p-2 w-full border rounded-md focus:border-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-300 transition-colors duration-300" />
                </div>
                <div>
                  <label htmlFor="password" className="block text-sm font-medium text-gray-700">Password</label>
                  <input ref={currentPasswordRef} type="password" id="password" name="password" className="mt-1 p-2 w-full border rounded-md focus:border-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-300 transition-colors duration-300" />
                </div>
                <div>
                  <label htmlFor="new-password" className="block text-sm font-medium text-gray-700">Confirm Password</label>
                  <input ref={newPasswordRef} type="password" id="new-password" name="new-password" className="mt-1 p-2 w-full border rounded-md focus:border-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-300 transition-colors duration-300" />
                </div>
                <div>
                  <label htmlFor="confirm-new-password" className="block text-sm font-medium text-gray-700">Confirm Password</label>
                  <input ref={confirmNewPasswordRef} type="password" id="confirm-new-password" name="confirm-new-password" className="mt-1 p-2 w-full border rounded-md focus:border-gray-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-300 transition-colors duration-300" />
                </div>
                <div className="w-full row-span-1 lg:h-1/4 space-x-3 flex justify-center ">
                  <div className="grid grid-cols-3 grid-rows-1 row-span-1 lg:h-1/4 space-x-3 pt-10">
                    <div>
                      <button onSubmit={deleteAccount} type="button" className="w-full bg-black text-white p-2 rounded-md hover:bg-gray-800 focus:outline-none focus:bg-black focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-900 transition-colors duration-300">Delete</button>
                    </div>
                    <div >
                      <button type="submit" className="w-full bg-black text-white p-2 rounded-md hover:bg-gray-800 focus:outline-none focus:bg-black focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-900 transition-colors duration-300">Sign Up</button>
                    </div>
                    <div >
                      <button onSubmit={logOut} type="button" className="w-full bg-black text-white p-2 rounded-md hover:bg-gray-800 focus:outline-none focus:bg-black focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-900 transition-colors duration-300">Logout</button>
                    </div>
                  </div>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div >
    </>
  )
}

export default EditProfileScreen





