package database

import (
	"fmt"
	"log"
	"os"

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
	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT")))
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	db.MustExec(indexing)

	return db
}
