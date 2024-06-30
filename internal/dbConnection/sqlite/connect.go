package sqlite

import (
	"database/sql"
	db "db-insert-app/internal/models"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteConn struct {
	// Connection to the database
	db *sql.DB
}

func New() *SQLiteConn {
	return &SQLiteConn{}
}

func (s *SQLiteConn) Connect(connStr string) (db.DB, error) {
	// Connect to the database
	db, err := sql.Open("sqlite3", connStr)

	if err != nil {
		return nil, err
	}

	s.db = db

	if db.Ping() != nil {
		return nil, err
	}

	return s, nil
}

func (s *SQLiteConn) Write() error {
	return nil
}

func (s *SQLiteConn) Close() error {
	return s.db.Close()
}
