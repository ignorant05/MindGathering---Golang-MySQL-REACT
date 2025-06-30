package handler

import (
	"backend/internal/db"
	"log/slog"
)

var Logger slog.Logger

type Handler struct {
	Queries *db.Queries
}
