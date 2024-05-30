package src

import (
	"encoding/json"
	"net/http"

	"github.com/mrspec7er/mailserv/repository/internal/dto"
)

type Controller struct {
	Service Service
}

func (c Controller) CreateEmail(w http.ResponseWriter, r *http.Request) {
	emails := []*dto.Email{}

	if err := json.NewDecoder(r.Body).Decode(&emails); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	status, err := c.Service.CreateEmail(emails)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{"data": emails})
}

func (c Controller) RetrieveEmail(w http.ResponseWriter, r *http.Request) {
	emails := []*dto.Email{}

	status, err := c.Service.RetrieveEmail(&emails)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]any{"data": emails})
}
