
import { useDispatch, useSelector } from "react-redux"
import { setBlogs, setPage } from "../redux/slices/pageSlice"
import { useEffect } from "react"
import BlogCard from "../components/BlogCard"
import { setInfo } from "../redux/slices/authSlice"
import { usePaginationQuery, useCountAllBlogsQuery } from "../redux/slices/usersApiSlice"

const HomeScreen = () => {
  const dispatch = useDispatch()

  const storedToken = sessionStorage.getItem('access_token');
  const auth = useSelector((state) => state.auth);
  const access_token = auth?.access_token || storedToken || "";

  useEffect(() => {
    setInfo({ access_token })
  }, [access_token])

  const blogsPerPage = 10
  var page = 0

  const { blogCount } = useCountAllBlogsQuery()
  const { blogs } = usePaginationQuery({ page: page, size: blogsPerPage })

  var totalPages = blogCount / blogsPerPage;

  const renderList = () => {
    if (blogs) {
      return blogs.map((blog) => <BlogCard key={blog.id} body={blog} />)
    }
    return (<span>No blogs</span>)
  }
  const getVisiblePages = (current, total) => {
    const delta = 1;
    const pages = [];

    for (let i = Math.max(1, current - delta); i <= Math.min(total, current + delta); i++) {
      pages.push(i);
    }

    return pages;
  };

  useEffect(() => {
    dispatch(setBlogs(blogs));
  }, [blogs, page, dispatch]);

  return (
    <>
      <div className="h-screen">
        <div className="max-w-5xl mx-auto mt-8">
          <div className="border-l-2 border-gray-500 pl-8">
            {
              renderList()
            }
          </div>
        </div>
        <div className="fixed bottom-0 left-0 w-full bg-white shadow-md p-10">
          <div className="flex items-center justify-center">
            <div className="py-3 border rounded-lg dark:border-gray-600">
              <ol
                className="flex items-center text-sm text-gray-500 divide-x rtl:divide-x-reverse divide-gray-300 dark:text-gray-400 dark:divide-gray-600">
                <li>
                  <button type="button"
                    onClick={() => { if (page > 1) dispatch(setPage(page - 1)) }}
                    className="relative flex items-center justify-center font-medium min-w-[2rem] px-1.5 h-8 -my-3 rounded-md outline-none hover:bg-gray-500/5 focus:bg-yellow-500/10 focus:ring-2 focus:ring-yellow-500 dark:hover:bg-gray-400/5 transition text-yellow-600"
                    aria-label="Previous" rel="prev">
                    <svg className="w-5 h-5 rtl:scale-x-[-1]" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"
                      fill="currentColor" aria-hidden="true">
                      <path fillRule="evenodd"
                        d="M12.707 5.293a1 1 0 010 1.414L9.414 10l3.293 3.293a1 1 0 01-1.414 1.414l-4-4a1 1 0 010-1.414l4-4a1 1 0 011.414 0z"
                        clipRule="evenodd"></path>
                    </svg>
                  </button>
                </li>

                {getVisiblePages(page, totalPages).map((p) => (
                  <li key={p}>
                    <button
                      onClick={() => dispatch(setPage(p))}
                      className={`px-2 h-8 rounded-md ${p === page ? "bg-yellow-500/10 ring-2 ring-yellow-500" : ""
                        }`}
                    >
                      {p}
                    </button>
                  </li>
                ))}
                <li>
                  <button type="button"
                    onClick={() => { if (page < totalPages) dispatch(setPage(page + 1)) }}
                    className="relative flex items-center justify-center font-medium min-w-[2rem] px-1.5 h-8 -my-3 rounded-md outline-none hover:bg-gray-500/5 focus:bg-yellow-500/10 focus:ring-2 focus:ring-yellow-500 dark:hover:bg-gray-400/5 transition text-yellow-600"
                    aria-label="Next" rel="next">
                    <svg className="w-5 h-5 rtl:scale-x-[-1]" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"
                      fill="currentColor" aria-hidden="true">
                      <path fillRule="evenodd"
                        d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
                        clipRule="evenodd"></path>
                    </svg>
                  </button>
                </li>
              </ol>
            </div>
          </div>
        </div>
      </div>
    </>
  )
}

export default HomeScreen
