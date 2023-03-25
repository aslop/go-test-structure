package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB interface {
	Connect() error
	Disconnect(ctx context.Context) error
	Collection(name string) *mongo.Collection
}

type MongoClient struct {
	Client *mongo.Client
}

func Connect() *mongo.Client {

	credentail := options.Credential{
		Username: os.Getenv("MONGO_DB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_DB_ROOT_PASSWORD"),
	}

	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(credentail)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		panic(err)
	}

	return client
}

func (c *MongoClient) Disconnect(ctx context.Context) error {
	return c.Client.Disconnect(ctx)
}

// Client instance
var Connection *mongo.Client = Connect()
var Database *mongo.Database = Connection.Database(os.Getenv("MONGO_DB_NAME"))

// getting database collections
func GetCollection(c *mongo.Client, collectionName string) *mongo.Collection {
	return c.Database(os.Getenv("MONGO_DB_NAME")).Collection(collectionName)
}
