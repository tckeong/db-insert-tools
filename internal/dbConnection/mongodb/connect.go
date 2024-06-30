package mongodb

import (
	"context"
	db "db-insert-app/internal/models"
	"fmt"
	"regexp"

	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDBConn is a struct that holds the connection to the MongoDB database
type MongoDBConn struct {
	// Connection to the database
	db             *mongo.Client
	collectionName string
	tableName      string
}

// New creates a new MongoDBConn instance
func New(dbName string) (*MongoDBConn, error) {
	regex := regexp.MustCompile(`\[(\w+)\/(\w+)\]`)

	match := regex.FindStringSubmatch(dbName)

	if len(match) > 2 {
		collectionName := match[1]
		tableName := match[2]
		return &MongoDBConn{collectionName: collectionName, tableName: tableName}, nil
	} else {
		return nil, fmt.Errorf("invalid database name")
	}
}

// Connect connects to the MongoDB database
func (m *MongoDBConn) Connect(connStr string) (db.DB, error) {
	dsn := fmt.Sprintf("%v/%v", connStr, m.tableName)

	// Set client options
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("error connecting to the database")
	}

	m.db = client

	return m, nil
}

// Write writes to the MongoDB database
func (m *MongoDBConn) Write() error {
	return nil
}

// Close closes the connection to the MongoDB database
func (m *MongoDBConn) Close() error {
	return m.db.Disconnect(context.TODO())
}
