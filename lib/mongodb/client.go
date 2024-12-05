package mongodb

import (
	"colegio/server/common/config"
	"colegio/server/common/utils"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client
var clientInstanceError error
var mongoOnce sync.Once

func NewClient(
	user string,
	password string,
	url string,
	dbName string,
	maxPool int,
	useSSL bool,
	timeout int64,
) (*mongo.Client, error) {
	connectionString := getConnectionStr(url, dbName, user, password)

	mongoOnce.Do(func() {
		//Set the context
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
		defer cancel()
		// Set client options
		clientOptions := options.Client().ApplyURI(connectionString)
		// Connect to MongoDB
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func NewClientDefault() (*mongo.Client, error) {
	dbConfig := config.GetConfigDefault().MongoDB
	dbClient, err := NewClient(
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.URL,
		dbConfig.DBName,
		dbConfig.MaxPool,
		dbConfig.UseSSL,
		dbConfig.Timeout,
	)
	return dbClient, errors.WithStack(err)
}

func getConnectionStr(
	url string,
	dbName string,
	user string,
	password string,
) string {
	if utils.GetStage() == utils.Local {
		return fmt.Sprintf("mongodb://%s", url)
	}
	return fmt.Sprintf("mongodb+srv://%s:%s@%s.mongodb.net/%s?retryWrites=true&w=majority",
		user,
		password,
		url,
		dbName,
	)
}

func GetCollection(collectionName string) *mongo.Collection {
	mongoClient, _ := NewClientDefault()
	dbConfig := config.GetConfigDefault().MongoDB
	return mongoClient.Database(dbConfig.DBName).Collection(collectionName)
}

func User() *mongo.Collection {
	return GetCollection("users")
}

func Task() *mongo.Collection {
	return GetCollection("tasks")
}

