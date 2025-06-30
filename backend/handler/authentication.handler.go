package handler

import (
	"backend/internal/db"
	"backend/internal/helpers"
	"backend/internal/utils"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload db.RegisterParams
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		Logger.Error("registration failed", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	password := helpers.HashPwd(payload.Password.String)

	formattedPwd := sql.NullString{
		String: password,
		Valid:  true,
	}

	payload.Password = formattedPwd

	err := h.Queries.Register(ctx, payload)
	if err != nil {
		Logger.Error("registration failed, user already exists", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	Logger.Info("registered")
}

type LoginStruct struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req LoginStruct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		Logger.Error("registration failed", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	payload := db.LoginParams{
		Username: req.Username,
		Email:    req.Email,
	}
	user, err := h.Queries.Login(ctx, payload)
	if err != nil {
		Logger.Error("User not found", "error", err, "path", r.URL.Path)
		http.Error(w, "Unauthorizaed, Access Denied", http.StatusUnauthorized)
		return
	}

	if !helpers.ComparePwd(user.Password.String, req.Password) {
		Logger.Error("Invalid username or password", "error", err, "path", r.URL.Path)
		http.Error(w, "Unauthorizaed, Access Denied", http.StatusUnauthorized)
		return
	}

	access_token, err := utils.GenerateAccessToken(user.Uid)
	if err != nil {
		Logger.Error("Error creating access token", "error", err, "path", r.URL.Path)
		http.Error(w, "Unauthorizaed, Access Denied", http.StatusUnauthorized)
		return
	}

	refresh_token, err := utils.GenerateRefreshToken(user.Uid)
	if err != nil {
		Logger.Error("Error creating refresh token", "error", err, "path", r.URL.Path)
		http.Error(w, "Unauthorizaed, Access Denied", http.StatusUnauthorized)
		return
	}

	w.Header().Add("Authorization", access_token)

	cookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refresh_token,
		MaxAge:   3600,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
	}

	http.SetCookie(w, &cookie)

	Logger.Info("Logged In")
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")
	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to ", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := h.Queries.Logout(ctx, user_id)

	if err != nil {
		Logger.Error("failed to logout", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found ", http.StatusUnauthorized)
		return
	}

	Logger.Info("Logged out")
}

type UserInfo struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")
	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to fetch user info", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user, err := h.Queries.GetUserById(ctx, user_id)
	if err != nil {
		Logger.Error("failed to fetch user data", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}
	res := UserInfo{
		ID:       user.Uid,
		Username: user.Username,
		Email:    user.Email,
	}

	Logger.Info("User credentials", "username", res.Username, "email", res.Email)
	json.NewEncoder(w).Encode(res)

}

func (h *Handler) UserProfile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userIdString := r.URL.Query().Get("userId")

	user_id, err := strconv.Atoi(userIdString)
	if err != nil {
		Logger.Error("failed to fetch user data", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	user, err := h.Queries.GetUserById(ctx, int64(user_id))
	if err != nil {
		Logger.Error("failed to fetch user data", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	res := UserInfo{
		ID:       user.Uid,
		Username: user.Username,
		Email:    user.Email,
	}

	Logger.Info("User credentials", "username", res.Username, "email", res.Email)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetMyPicture(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")
	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to fetch user info", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	payload := sql.NullInt64{
		Int64: int64(user_id),
		Valid: true,
	}

	picture, err := h.Queries.GetUserProfilePic(ctx, payload)
	if err != nil {
		Logger.Error("failed to fetch picture", "error", err, "path", r.URL.Path)
		http.Error(w, "not found not found", http.StatusUnauthorized)
		return
	}

	picBase64 := base64.StdEncoding.EncodeToString(picture.Data)
	if picBase64 == "" {
		Logger.Error("failed to convert image", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	imageDataURL := fmt.Sprintf("data:image/%s;base64,%s", picture.Type, picBase64)

	Logger.Info("User credentials", "imageDataUrl", imageDataURL)
	json.NewEncoder(w).Encode(imageDataURL)
}

func (h *Handler) GetUserPicture(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID := chi.URLParam(r, "userid")
	user_id, err := strconv.Atoi(userID)
	if userID == "" || err != nil {
		Logger.Error("Not a valid id", "error", err, "path", r.URL.Path)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = h.Queries.GetUserById(ctx, int64(user_id))
	if err != nil {
		Logger.Error("failed to fetch user data", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	payload := sql.NullInt64{
		Int64: int64(user_id),
		Valid: true,
	}

	picture, err := h.Queries.GetUserProfilePic(ctx, payload)
	if err != nil {
		Logger.Error("failed to fetch picture", "error", err, "path", r.URL.Path)
		http.Error(w, "not found not found", http.StatusUnauthorized)
		return
	}

	picBase64 := base64.StdEncoding.EncodeToString(picture.Data)
	if picBase64 == "" {
		Logger.Error("failed to convert image", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	imageDataURL := fmt.Sprintf("data:image/%s;base64,%s", picture.Type, picBase64)

	Logger.Info("User credentials", "imageDataUrl", imageDataURL)
	json.NewEncoder(w).Encode(imageDataURL)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")
	uid, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to fetch user info", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user, err := h.Queries.GetUserById(ctx, uid)
	if err != nil {
		Logger.Error("user not found", "error", err, "path", r.URL.Path)
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	err = r.ParseMultipartForm(10 << 20) // Max 10MB
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	oldPassword := r.FormValue("oldPassword")
	newPassword := r.FormValue("newPassword")

	file, handler, err := r.FormFile("profilePic")
	if err != nil && err != http.ErrMissingFile {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		return
	}

	var imageData []byte
	if file != nil {
		defer file.Close()
		imageData, err = io.ReadAll(file)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}

	if !helpers.ComparePwd(user.Password.String, oldPassword) {
		Logger.Error("Invalid username or password", "path", r.URL.Path)
		http.Error(w, "Unauthorizaed, Access Denied", http.StatusUnauthorized)
		return
	}

	var pwd string
	if newPassword == "" {
		pwd = oldPassword
	} else {
		pwd = newPassword
	}
	password := sql.NullString{
		String: helpers.HashPwd(pwd),
		Valid:  true,
	}

	draft := db.UpdateUserParams{
		Username: username,
		Email:    email,
		Password: password,
	}

	err = h.Queries.UpdateUser(ctx, draft)
	if err != nil {
		Logger.Error("failed to update credentials", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found ", http.StatusUnauthorized)
		return
	}

	payload := db.UpdateUserPicParams{
		Name: handler.Filename,
		Type: handler.Header.Get("Content-Type"),
		Data: imageData,
	}

	err = h.Queries.UpdateUserPic(ctx, payload)
	if err != nil {
		Logger.Error("failed to update credentials", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found ", http.StatusUnauthorized)
		return
	}

	Logger.Info("Updated")
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	uidVal := r.Context().Value("user_id")
	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to delete user", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err := h.Queries.DeleteUser(ctx, user_id)

	if err != nil {
		Logger.Error("failed to logout", "error", err, "path", r.URL.Path)
		http.Error(w, "User not found ", http.StatusUnauthorized)
		return
	}

	Logger.Info("User Deleted")
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	uidVal := r.Context().Value("user_id")

	user_id, ok := uidVal.(int64)
	if !ok {
		Logger.Error("failed to delete user", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	token, err := h.Queries.GetTokenByUid(ctx, user_id)
	if err != nil {
		Logger.Error("failed to fetch refresh token ", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, err = utils.ValidateRefreshToken(token.Token)
	if err != nil {
		Logger.Error("refresh token not valid", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	access_token, err := utils.GenerateAccessToken(user_id)
	if err != nil {
		Logger.Error("failed Create new access token", "path", r.URL.Path)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Authorization", access_token)
	Logger.Info("Refreshed")
}
