package routes

import (
	"backend/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"backend/internal/db"
	middlewares "backend/internal/middleware"
)

func LoadRoutes(queries *db.Queries) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/api/v1/users", func(r chi.Router) {
		loadAuthRoutes(r, queries)
	})

	router.Route("/api/v1/auth", func(r chi.Router) {
		r.Use(middlewares.VerifyAccessTokenMiddleware)
		loadUserRoutes(r, queries)
	})

	router.Route("/api/v1/auth/refresh", func(r chi.Router) {
		r.Use(middlewares.VerifyAccessTokenMiddleware)
		r.Use(middlewares.VerifyRefreshTokenMiddleware)
		loadRefreshRoutes(r, queries)
	})
	return router
}

func loadAuthRoutes(router chi.Router, queries *db.Queries) {
	handler := &handler.Handler{Queries: queries}

	router.Post("/register", handler.Register)
	router.Post("/login", handler.Login)
}

func loadUserRoutes(router chi.Router, queries *db.Queries) {
	handler := &handler.Handler{Queries: queries}

	router.Delete("/delete", handler.Delete)
	router.Post("/logout", handler.Logout)
	router.Get("/profile", handler.Profile)
	router.Get("/profile/info", handler.UserProfile)
	router.Put("/update/credentials", handler.Update)
	router.Get("/get/my/image", handler.GetMyPicture)
	router.Get("/get/users/{userId}/image", handler.GetUserPicture)

	router.Get("/pages", handler.Pagination)

	router.Get("/get/blogs", handler.GetAllBlogs)
	router.Get("/get/blog", handler.GetBlog)
	router.Get("/get/user/blogs", handler.GetUserBlog)
	router.Post("/create/blog", handler.CreateNewBlog)
	router.Put("/update/blogs", handler.EditBlog)
	router.Delete("/delete/blogs", handler.DeleteBlog)

	router.Get("/get/comments/users", handler.GetAuthorName)
	router.Get("/get/blogs/{blogId}/comments", handler.GetCommentsByBId)
	router.Get("/get/users/{userId}/comments", handler.GetCommentsByAId)
	router.Post("/create/blog/{blogId}/comment", handler.CommentOnBlog)
	router.Put("/update/blogs/{blogId}/comments", handler.EditComment)
	router.Delete("/delete/blogs/{blogId}/comments", handler.DeleteComment)

	router.Get("/count/users/{userId}/blogs", handler.CountUserBlogs)
	router.Get("/count/all/blogs", handler.CountBlogs)
	router.Get("/count/my/blogs/", handler.CountMyBlogs)
	router.Get("/count/blogs/{blogId}/comments", handler.CountCommentsForBlog)
	router.Get("/count/all/comments", handler.CountComments)
	router.Get("/count/my/comments", handler.CountMyComments)
	router.Get("/count/users/{userId}/comments", handler.CountUserComments)
}

func loadRefreshRoutes(router chi.Router, queries *db.Queries) {
	handler := &handler.Handler{Queries: queries}
	router.Post("/", handler.Refresh)
}
