package database

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS emails (
    id SERIAL PRIMARY KEY,
    subject TEXT,
    sender TEXT,
    recipient TEXT,
    date TIMESTAMP,
    body TEXT,
	flag TEXT DEFAULT 'pending'
);`

var indexing = `
CREATE INDEX IF NOT EXISTS idx_email_flag ON emails (flag)
`

func StartConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres password=mrc201 dbname=coldsiretest host=localhost sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	db.MustExec(indexing)

	return db
}
