package dto

import "time"

type Email struct {
	ID        uint      `db:"id" json:"id"`
	Subject   string    `db:"subject" json:"subject"`
	Sender    string    `db:"sender" json:"sender"`
	Recipient string    `db:"recipient" json:"recipient"`
	Date      time.Time `db:"date" json:"date"`
	Body      string    `db:"body" json:"body"`
	Flag      string    `db:"flag" json:"flag"`
}
