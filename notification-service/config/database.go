package config

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
	CREATE TABLE IF NOT EXISTS notifications (
		id SERIAL PRIMARY KEY,
		email TEXT NOT NULL,
		message TEXT NOT NULL,
		type TEXT NOT NULL,
		is_send BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	)
`

func InitDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=junction.proxy.rlwy.net port=36991 user=postgres password=XNincXXJrBYOanQFJxfBHMmaRRjgQoqY dbname=railway sslmode=disable")
	if err != nil {
		return nil, err
	}

	db.MustExec(schema)

	return db, nil
}
