package middlewares

import (
	"backend/internal/utils"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var logger = slog.Default()

func VerifyAccessTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			logger.Error("Empty header", "path", r.URL.Path)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		parsedToken := authHeader[len("Bearer ")+1:]

		token, err := utils.ValidateAccessToken(parsedToken)
		if err != nil {
			logger.Error("Invalid Access Token", "error", err, "path", r.URL.Path)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Error("Invalid Access Token", "error", err, "path", r.URL.Path)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		expFloat, ok := claims["exp"].(float64)
		if !ok {
			logger.Error("Expired", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}
		exp := int64(expFloat)
		if exp < time.Now().Unix() {
			logger.Error("Expired", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(float64)
		fmt.Printf("sub: %v", sub)
		if !ok {
			logger.Error("User not found", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		userID := sql.NullInt64{
			Int64: int64(sub),
			Valid: true,
		}

		ctx := context.WithValue(r.Context(), "user_id", userID)
		r = r.WithContext(ctx)

		logger.Info("Token verified")
		next.ServeHTTP(w, r)
	})
}

func VerifyRefreshTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parsedToken, err := r.Cookie("refresh_token")
		if err != nil {
			logger.Error("No coookie", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		token, err := utils.ValidateRefreshToken(parsedToken.Value)
		if err != nil {
			logger.Error("Invalid refresh token", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			logger.Error("Invalid refresh token", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		expFloat, ok := claims["exp"].(float64)
		if !ok {
			logger.Error("Invalid refresh token", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		exp := int64(expFloat)
		if exp < time.Now().Unix() {
			logger.Error("Invalid refresh token", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			logger.Error("No user found", "error", err, "path", r.URL.Path)
			http.Error(w, "Unauthorized, Access Denied", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", int64(sub))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)

	})
}
