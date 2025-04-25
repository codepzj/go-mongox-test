package global

import (
	"context"
	"github.com/chenmingyong0423/go-mongox/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

// NewMongoClient mongo客户端实例
func NewMongoClient() (*mongo.Client, error) {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{
		Username:   "test",
		Password:   "test",
		AuthSource: "db-test",
	}))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		panic(err)
	}
	return client, nil
}

func NewMongoDatabase() *mongox.Database {
	mongoClient, err := NewMongoClient()
	if err != nil {
		panic(err)
	}
	return mongox.NewClient(mongoClient, &mongox.Config{}).NewDatabase("db-test")
}

var DB *mongox.Database

func init() {
	DB = NewMongoDatabase()
}
