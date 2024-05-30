package src

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mrspec7er/mailserv/repository/internal/dto"
)

type Service struct {
	DB *sqlx.DB
}

func (s Service) CreateEmail(email *dto.Email) (int, error) {
	const insertQuery = `INSERT INTO emails (subject, sender, recipient, date, body) VALUES ($1, $2, $3, $4, $5)`

	stmt, err := s.DB.Prepare(insertQuery)
	if err != nil {
		return 400, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(email.Subject, email.Sender, email.Recipient, time.Now(), email.Body)
	if err != nil {
		log.Fatal("error inserting email: %w", err)
	}

	return 201, nil

}

func (s Service) RetrieveEmail(emails *[]*dto.Email) (int, error) {
	query := `UPDATE emails SET flag = $1 WHERE flag = $2 RETURNING id, subject, sender, recipient, date, body, flag`

	err := s.DB.Select(emails, query, "read", "pending")
	if err != nil {
		return 400, err
	}

	return 200, nil
}
