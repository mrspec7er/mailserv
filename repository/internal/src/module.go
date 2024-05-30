package src

import (
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Module(DB *sqlx.DB) func(chi.Router) {

	c := Controller{
		Service: Service{
			DB: DB,
		},
	}

	return func(r chi.Router) {
		r.Post("/", c.CreateEmail)
		r.Get("/", c.RetrieveEmail)
	}
}
