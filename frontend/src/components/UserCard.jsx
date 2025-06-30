const UserCard = (props) => {
  const { id, title, author, updated_at } = props.body
  const dispatch = useDispatch()
  const navigate = useNavigate()


  const viewBlogPost = () => {
    navigate(`/get/users?${userId}`)
  }

  const deleteUser = async () => {
    fetch(`/users/delete/blogs?blogId=${id}`).then(() => { console.log("deleted") }).catch((err) => { console.error(err) })
    dispatch(clearClickedBlog())
  }

  return (
    <>
      <div className="flex flex-col md:flex-row md:justify-between mt-8">
        <Link to={`/users/blogs/${id}`} className="mb-4 md:mb-0">
          <h3 className="text-xl font-bold mb-2">{title}</h3>
          <p className="text-gray-600 text-sm">{updated_at}</p>
        </Link>
        <p className="text-gray-700">Created by :{author}</p>
        <Menu as="div" className="relative inline-block text-left">
          <div>
            <MenuButton className="inline-flex w-full justify-center gap-x-1.5 rounded-md bg-white px-3 py-2 text-sm font-semibold text-gray-900 shadow-xs ring-1 ring-gray-300 ring-inset hover:bg-gray-50">
              Options
              <ChevronDownIcon aria-hidden="true" className="-mr-1 size-5 text-gray-400" />
            </MenuButton>
          </div>

          <MenuItems
            transition
            className="absolute right-0 z-10 mt-2 w-56 origin-top-right rounded-md bg-white shadow-lg ring-1 ring-black/5 transition focus:outline-hidden data-closed:scale-95 data-closed:transform data-closed:opacity-0 data-enter:duration-100 data-enter:ease-out data-leave:duration-75 data-leave:ease-in"
          >
            <div className="py-1">
              <MenuItem>
                <button
                  onClick={editBlogPost}
                  className="block px-4 py-2 text-sm text-gray-700 data-focus:bg-gray-100 data-focus:text-gray-900 data-focus:outline-hidden"
                >
                  Edit
                </button>
              </MenuItem>
              <MenuItem>
                <button
                  onClick={deleteBlogPost}
                  className="block px-4 py-2 text-sm text-gray-700 data-focus:bg-gray-100 data-focus:text-gray-900 data-focus:outline-hidden"
                >
                  Delete
                </button>
              </MenuItem>
            </div>
          </MenuItems>
        </Menu>
      </div >
    </>
  )
}

export default UserCard
