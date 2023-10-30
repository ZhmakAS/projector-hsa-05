package pkg

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

func InitMySQL(ctx context.Context, uri string) (*sqlx.DB, error) {
	connect, err := sqlx.Open("mysql", uri)
	if err != nil {
		log.Printf("Failed to open mysql connection: %s", err)
		return nil, err
	}

	if err := connect.Ping(); err != nil {
		log.Printf("Could not ping mysql: %s", err)
		return nil, err
	}

	return connect, nil
}
