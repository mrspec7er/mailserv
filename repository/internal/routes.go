package internal

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrspec7er/mailserv/repository/internal/src"
)

func (s Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello There!"})
	})

	router.Route("/emails", src.Module(s.DB))

	return router
}
