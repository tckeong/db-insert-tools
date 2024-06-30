package mysql

import (
	"database/sql"
	db "db-insert-app/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConn struct {
	// Connection to the database
	db *sql.DB
}

func New() *MySQLConn {
	return &MySQLConn{}
}

func (m *MySQLConn) Connect(connStr string) (db.DB, error) {
	// Connect to the database
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		return nil, err
	}

	m.db = db

	if db.Ping() != nil {
		return nil, err
	}

	return m, nil
}

func (m *MySQLConn) Write() error {
	return nil
}

func (m *MySQLConn) Close() error {
	return m.db.Close()
}
