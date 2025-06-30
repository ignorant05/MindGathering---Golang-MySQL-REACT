import { useEffect } from "react";
import { setInfo } from "../../redux/slices/authSlice";
import { useSelector } from "react-redux";
import { Link } from "react-router-dom";

const Navbar = () => {

  const storedToken = sessionStorage.getItem('access_token');
  const storedUser = sessionStorage.getItem('user');
  const access_token = useSelector((state) => state.auth.access_token) || storedToken || "";
  const user = useSelector((state) => state.auth.user) || storedUser || "";

  useEffect(() => {
    setInfo({ user, access_token })
  }, [access_token, user])

  const loggedIn = () => {
    if (access_token) {
      return <div className="flex items-center justify-end gap-3">
        <Link className="hidden items-center justify-center rounded-xl bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 transition-all duration-150 hover:bg-gray-50 sm:inline-flex"
          to="/profile">{user.username}</Link>
      </div>
    }
    return <div className="flex items-center justify-end gap-3">
      <Link className="hidden items-center justify-center rounded-xl bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 transition-all duration-150 hover:bg-gray-50 sm:inline-flex"
        to="/register">Sign in</Link>
      <Link className="inline-flex items-center justify-center rounded-xl bg-blue-600 px-3 py-2 text-sm font-semibold text-white shadow-sm transition-all duration-150 hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600"
        to="/login">Login</Link>
    </div>
  }

  return (
    <>
      <header
        className="fixed inset-x-0 top-0 z-30 mx-auto w-full max-w-screen-md border border-gray-100 bg-white/80 py-3 shadow backdrop-blur-lg md:top-6 md:rounded-3xl lg:max-w-screen-lg">
        <div className="px-4">
          <div className="flex items-center justify-between">
            <div className="flex shrink-0">
              <Link aria-current="page" className="flex items-center" to={"/"}>
                <img className="h-7 w-auto" src="../../assets/logo.jpg" alt="" />
                <p className="sr-only">Website Title</p>
              </Link>
            </div>
            <div className="hidden md:flex md:items-center md:justify-center md:gap-5">
              <Link aria-current="page"
                className="inline-block rounded-lg px-2 py-1 text-sm font-medium text-gray-900 transition-all duration-200 hover:bg-gray-100 hover:text-gray-900"
                to={"/"}>Home</Link>
            </div>
            {loggedIn()}
          </div>
        </div>
      </header>
    </>
  )
}

export default Navbar
