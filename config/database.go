package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
	CREATE TABLE notifications (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL,
		message TEXT NOT NULL,
		type TEXT NOT NULL,
		is_send BOOLEAN NOT NULL
	)
`

func InitDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres password=12345678 dbname=jec-live-code sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.MustExec(schema)

	return db, nil
}
