package server

import (
	config "backend/config/db"
	"backend/internal/db"
	"backend/internal/helpers"
	"backend/internal/routes"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"log/slog"

	"context"
)

var logger slog.Logger

type App struct {
	DB      *sql.DB
	Router  http.Handler
	Queries *db.Queries
}

func New() *App {
	conn, err := config.ConnMySql()
	if err != nil {
		logger.Error("Failed to connect to database", "error", err)
		return nil
	}

	queries := db.New(conn)
	app := &App{
		DB:      conn,
		Queries: queries,
		Router:  routes.LoadRoutes(queries),
	}

	return app
}

func (a *App) Run(ctx context.Context) error {
	helpers.LoadEnvironmentFile()
	PORT := os.Getenv("SERVER_PORT")
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: a.Router,
	}
	var err error

	defer func() {
		if err := a.DB.Close(); err != nil {
			logger.Error("Failed to close db connection", "error", err)
		}
	}()

	ch := make(chan error, 1)
	go func() {
		err = server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start server %w", err)
		}
		close(ch)
	}()

	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(timeout)
	}

}
