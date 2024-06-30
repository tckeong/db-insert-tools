package postgresql

import (
	"database/sql"
	"errors"

	db "db-insert-app/internal/models"

	_ "github.com/lib/pq"
)

type PostgreSQLConn struct {
	// Connection to the database
	db *sql.DB
}

func New() *PostgreSQLConn {
	return &PostgreSQLConn{}
}

func (p *PostgreSQLConn) Connect(connStr string) (db.DB, error) {
	// Connect to the database
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, errors.New("error connecting to the database")
	}

	p.db = db

	if db.Ping() != nil {
		return nil, errors.New("error pinging the database")
	}

	return p, nil
}

func (p *PostgreSQLConn) Write() error {
	return nil
}

func (p *PostgreSQLConn) Close() error {
	return p.db.Close()
}
