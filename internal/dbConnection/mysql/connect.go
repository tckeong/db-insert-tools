package mysql

import (
	"database/sql"
	"db-insert-app/internal/models"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLConn struct {
	// Connection to the database
	db        *sql.DB
	tableName string
}

func New(tableName string) *MySQLConn {
	return &MySQLConn{
		tableName: tableName,
	}
}

func (m *MySQLConn) Connect(connStr string) (models.DB, error) {
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

func (m *MySQLConn) Write(data []models.Pair) error {
	db := m.db
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

	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", m.tableName, queryKeys, queryPlaceholders)

	_, err := db.Exec(query, values...)
	return err
}

func (m *MySQLConn) Close() error {
	return m.db.Close()
}
