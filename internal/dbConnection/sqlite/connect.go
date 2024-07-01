package sqlite

import (
	"database/sql"
	"db-insert-app/internal/models"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteConn struct {
	// Connection to the database
	db        *sql.DB
	tableName string
}

func New(tableName string) *SQLiteConn {
	return &SQLiteConn{
		tableName: tableName,
	}
}

func (s *SQLiteConn) Connect(dsn string) (models.DB, error) {
	// Connect to the database
	db, err := sql.Open("sqlite3", dsn)

	if err != nil {
		return nil, err
	}

	s.db = db

	if db.Ping() != nil {
		return nil, err
	}

	return s, nil
}

func (s *SQLiteConn) Write(data []models.Pair) error {
	db := s.db
	keys := make([]string, len(data))
	placeholders := make([]string, len(data))
	values := make([]interface{}, len(data))

	for i, pair := range data {
		keys[i] = pair.Key
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		values[i] = pair.Value
	}

	queryKeys := strings.Join(keys, ", ")
	queryPlaceholders := strings.Join(placeholders, ", ")

	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", s.tableName, queryKeys, queryPlaceholders)

	_, err := db.Exec(query, values...)
	return err
}

func (s *SQLiteConn) Close() error {
	return s.db.Close()
}
