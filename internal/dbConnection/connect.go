package db

import (
	"fmt"

	MongoDB "db-insert-app/internal/dbConnection/mongodb"
	MySQL "db-insert-app/internal/dbConnection/mysql"
	PostgreSQL "db-insert-app/internal/dbConnection/postgresql"
	SQLite "db-insert-app/internal/dbConnection/sqlite"
	db "db-insert-app/internal/models"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	db       db.DB
}

func New(host, port, username, password, dbname string, dbType int) *DBConnection {
	dbConn := &DBConnection{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbname,
	}

	dsn := dbConn.getConnStr(dbType)

	switch dbType {
	case 0:
		db := PostgreSQL.New()
		db.Connect(dsn)
		dbConn.db = db
	case 1:
		db := MySQL.New()
		db.Connect(dsn)
		dbConn.db = db
	case 2:
		db := SQLite.New()
		db.Connect(dsn)
		dbConn.db = db
	case 3:
		db, _ := MongoDB.New(dbname)
		db.Connect(dsn)
		dbConn.db = db
	}

	return dbConn
}

func (DB *DBConnection) getConnStr(dbType int) string {
	switch dbType {
	case 0:
		return fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable", DB.Host, DB.Port, DB.Username, DB.Password, DB.DBName)
	case 1:
		return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", DB.Username, DB.Password, DB.Host, DB.Port, DB.DBName)
	case 2:
		return fmt.Sprintf("%v", DB.DBName)
	case 3:
		return fmt.Sprintf("mongodb://%v:%v@%v:%v", DB.Username, DB.Password, DB.Host, DB.Port)
	default:
		// unreachable
		return ""
	}
}

func (DB *DBConnection) Close() error {
	return DB.db.Close()
}

func (DB *DBConnection) Write() error {
	return nil
}
