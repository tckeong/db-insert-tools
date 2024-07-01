package mongodb

import (
	"context"
	"db-insert-app/internal/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
func New(tableName, dbName string) *MongoDBConn {
	return &MongoDBConn{collectionName: dbName, tableName: tableName}
}

// Connect connects to the MongoDB database
func (m *MongoDBConn) Connect(dsn string) (models.DB, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("error connecting to the database")
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	m.db = client

	return m, nil
}

// Write writes to the MongoDB database
func (m *MongoDBConn) Write(data []models.Pair) error {
	// Get a handle for your collection
	collection := m.db.Database(m.tableName).Collection(m.collectionName)

	// Create a document to insert
	doc := bson.D{}

	for _, pair := range data {
		doc = append(doc, bson.E{Key: pair.Key, Value: pair.Value})
	}

	// Insert the document into the collection
	_, err := collection.InsertOne(context.TODO(), doc)
	return err
}

// Close closes the connection to the MongoDB database
func (m *MongoDBConn) Close() error {
	return m.db.Disconnect(context.TODO())
}
