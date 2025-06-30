import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import RegistrationScreen from './screens/RegistrationScreen'
import LoginScreen from './screens/LoginScreen'
import HomeScreen from './screens/HomeScreen'
import BlogViewScreen from './screens/BlogViewScreen'
import Navbar from './components/shared/Navbar'
import CreateBlogScreen from './screens/CreateBlogScreen'
import MyBlogsListScreen from './screens/MyBlogsListScreen'
import EditBlogScreen from './screens/EditBlogScreen'
import EditProfileScreen from './screens/EditProfileScreen'
import SeeProfileScreen from './screens/SeeProfileScreen'
import MyCommentsListScreen from './screens/MyCommentsListScreen'
import SeeOtherUserProfileScreen from './screens/SeeOtherUserProfileScreen'

function AppLayout() {
  return (
    <>
      <Navbar />
      <main className="pt-[72px]">
        <Routes>
          <Route path="/" element={<HomeScreen />} />
          <Route path="/home" element={<HomeScreen />} />
          <Route path="/my/blogs" element={<MyBlogsListScreen />} />
          <Route path="/my/comments" element={<MyCommentsListScreen />} />
          <Route path="/get/blog" element={<BlogViewScreen />} />
          <Route path="/edit/blog" element={<EditBlogScreen />} />
          <Route path="/create/blog" element={<CreateBlogScreen />} />
          <Route path="/edit/profile" element={<EditProfileScreen />} />
          <Route path="/users/profile" element={<SeeOtherUserProfileScreen />} />
        </Routes>
      </main>
    </>
  );
}

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/register" element={<RegistrationScreen />} />
        <Route path="/login" element={<LoginScreen />} />
        <Route path="/profile" element={<SeeProfileScreen />} />
        <Route path="*" element={<AppLayout />} />
      </Routes>
    </BrowserRouter>
  );
}
export default App
