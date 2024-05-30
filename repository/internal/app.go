package internal

import (
	"net/http"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rs/cors"
)

type Server struct {
	DB *sqlx.DB
}

func NewServer(s Server) *http.Server {
	c := cors.Default()
	server := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      c.Handler(s.RegisterRoutes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
