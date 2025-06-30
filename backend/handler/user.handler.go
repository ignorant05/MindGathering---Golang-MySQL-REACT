package handler

import (
	"backend/internal/db"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) Pagination(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageStr := r.URL.Query().Get("page")
	sizeStr := r.URL.Query().Get("size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10
	}
	offset := (page - 1) * size
	pageParams := db.PaginationParams{
		Offset: int32(offset),
		Limit:  int32(size),
	}
	blogs, err := h.Queries.Pagination(ctx, pageParams)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}
	totalCount := len(blogs) / size
	Logger.Info("success", "blogs", blogs)
	w.Header().Add("X-Total-Count", strconv.Itoa(totalCount))
	json.NewEncoder(w).Encode(blogs)
}

func (h *Handler) CountBlogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogsNumber, err := h.Queries.CountAllBlogs(ctx)
	if err != nil {
		Logger.Error("something went wrong", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogsNumber)
}

func (h *Handler) CountComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	commentsNumber, err := h.Queries.CountAllComments(ctx)
	if err != nil {
		Logger.Error("something went wrong", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(commentsNumber)
}

func (h *Handler) CountMyComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")

	userId, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to count comments", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	commentsNumber, err := h.Queries.CountMyComments(ctx, int64(userId))
	if err != nil {
		Logger.Error("something went wrong", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(commentsNumber)
}

func (h *Handler) CountUserComments(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	userIdStr := r.URL.Query().Get("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	commentsNumber, err := h.Queries.CountMyComments(ctx, int64(userId))
	if err != nil {
		Logger.Error("something went wrong", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(commentsNumber)
}

func (h *Handler) CountUserBlogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdStr := r.URL.Query().Get("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	blogsNumber, err := h.Queries.CountMyBlogs(ctx, int64(userId))
	if err != nil {
		Logger.Error("something went wrong", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogsNumber)
}

func (h *Handler) CountMyBlogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")

	userId, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to count blogs", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	blogsNumber, err := h.Queries.CountMyBlogs(ctx, int64(userId))
	if err != nil {
		Logger.Error("something went wrong", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(blogsNumber)
}

func (h *Handler) CountCommentsForBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogID := r.URL.Query().Get("blogId")
	blog_id, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	commentsNumber, err := h.Queries.CountMyComments(ctx, int64(blog_id))
	if err != nil {
		Logger.Error("something went wrong", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(commentsNumber)
}

func (h *Handler) GetAuthorName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdStr := r.URL.Query().Get("userId")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	username, err := h.Queries.GetAuthorName(ctx, int64(userId))
	if err != nil {
		http.Error(w, "Query error", http.StatusInternalServerError)
		return
	}

	Logger.Info("success", "username", username)
	json.NewEncoder(w).Encode(username)
}

func (h *Handler) GetBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	blogID := r.URL.Query().Get("blogId")
	blog_id, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	blog, err := h.Queries.GetBlogByBId(ctx, int64(blog_id))
	if err != nil {
		Logger.Error("No blog found by this id", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("success", "blog", blog)
	json.NewEncoder(w).Encode(blog)
}

func (h *Handler) GetUserBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to fetch blogs", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	blogs, err := h.Queries.GetUserBlogs(ctx, user_id)
	if err != nil {
		Logger.Error("No blog found by this id", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("success", "blogs", blogs)
	json.NewEncoder(w).Encode(blogs)
}

func (h *Handler) GetAllBlogs(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogs, err := h.Queries.GetAllBlogs(ctx)
	if err != nil {
		Logger.Error("No blog found", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("success", "blogs", blogs)
	json.NewEncoder(w).Encode(blogs)
}

func (h *Handler) CreateNewBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload db.NewBlogParams
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		Logger.Error("Cannot create blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := h.Queries.NewBlog(ctx, payload)
	if err != nil {
		Logger.Error("Cannot create blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("Success")
}

type BlogOpsParams struct {
	Title string
	Bid   int64
}

type EditBlogParams struct {
	Title   string
	Content string
}

func (h *Handler) EditBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogID := r.URL.Query().Get("blogId")
	blog_id, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	var payload EditBlogParams
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		Logger.Error("Cannot update blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to update blog", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	params := db.UpdateBlogByBIdParams{
		Bid:      int64(blog_id),
		AuthorID: user_id,
		Title:    payload.Title,
		Content:  payload.Content,
	}
	err = h.Queries.UpdateBlogByBId(ctx, params)
	if err != nil {
		Logger.Error("Cannot update this blog post", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("Success")
}

func (h *Handler) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogID := r.URL.Query().Get("blogId")
	blog_id, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to delete blog", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	params := db.DeleteBlogParams{
		Bid:      int64(blog_id),
		AuthorID: user_id,
	}
	err = h.Queries.DeleteBlog(ctx, params)
	if err != nil {
		Logger.Error("Cannot delete this blog post", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("Success")
}

func (h *Handler) CommentOnBlog(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogID := chi.URLParam(r, "blogid")
	blog_id, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var payload db.NewCommentParams
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		Logger.Error("Cannot comment on this blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to delete user", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	payload.AuthorID = user_id
	payload.BlogID = int64(blog_id)

	err = h.Queries.NewComment(ctx, payload)
	if err != nil {
		Logger.Error("Cannot comment on this blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("Success")
}

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	comments, err := h.Queries.GetAllComments(ctx)
	if err != nil {
		Logger.Error("No comments found", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("success", "comments", comments)
	json.NewEncoder(w).Encode(comments)
}

func (h *Handler) GetCommentsByBId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogID := chi.URLParam(r, "blogid")
	bid, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	comments, err := h.Queries.GetCommentsForBlogByBId(ctx, int64(bid))
	if err != nil {
		Logger.Error("No comments found", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("success", "comments", comments)
	json.NewEncoder(w).Encode(comments)
}

func (h *Handler) GetCommentsByAId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to fetch comments", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	comments, err := h.Queries.GetCommentsForUserByAid(ctx, user_id)
	if err != nil {
		Logger.Error("No comments found", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("success", "comments", comments)
	json.NewEncoder(w).Encode(comments)
}

type CommentOpsParams struct {
	Cid     int64
	Content string
}

func (h *Handler) EditComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogID := chi.URLParam(r, "blogid")
	bid, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	var payload CommentOpsParams
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		Logger.Error("Cannot comment on this blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to delete user", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	params := db.UpdateCommentParams{
		BlogID:   int64(bid),
		Cid:      payload.Cid,
		AuthorID: user_id,
		Content:  payload.Content,
	}

	err = h.Queries.UpdateComment(ctx, params)
	if err != nil {
		Logger.Error("Cannot comment on this blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("Success")
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	blogID := chi.URLParam(r, "blogid")
	bid, err := strconv.Atoi(blogID)
	if blogID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	commentID := chi.URLParam(r, "commentid")
	cid, err := strconv.Atoi(commentID)
	if commentID == "" || err != nil {
		Logger.Error("Not a valid blog", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to delete comment", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	payload := db.DeleteCommentParams{
		Cid:      int64(cid),
		BlogID:   int64(bid),
		AuthorID: user_id,
	}
	err = h.Queries.DeleteComment(ctx, payload)
	if err != nil {
		Logger.Error("Cannot delete this comment", "error", err, "path", r.URL.Path)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	Logger.Info("Success")
}
