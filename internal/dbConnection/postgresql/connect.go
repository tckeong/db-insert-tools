package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"db-insert-app/internal/models"

	_ "github.com/lib/pq"
)

type PostgreSQLConn struct {
	// Connection to the database
	db        *sql.DB
	tableName string
}

func New(tableName string) *PostgreSQLConn {
	return &PostgreSQLConn{
		tableName: tableName,
	}
}

func (p *PostgreSQLConn) Connect(connStr string) (models.DB, error) {
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

func (p *PostgreSQLConn) Write(data []models.Pair) error {
	db := p.db
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

	query := fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", p.tableName, queryKeys, queryPlaceholders)

	_, err := db.Exec(query, values...)
	return err
}

func (p *PostgreSQLConn) Close() error {
	return p.db.Close()
}
