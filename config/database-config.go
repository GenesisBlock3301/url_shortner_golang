package config

import (
	"context"
	"fmt"
	"github.com/GenesisBlock3301/url_shortner_golang/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	mongoDB *mongo.Database
	client  *mongo.Client
	err     error
)

const (
	connectTimeout = 30
)

func Connect() {
	databaseName := DatabaseName
	logger.Log{Message: fmt.Sprintf("Mongo url: %v", MongoUrl)}.Info()
	client, err := mongo.NewClient(options.Client().ApplyURI(MongoUrl))
	if err != nil {
		logger.Log{Message: fmt.Sprintf("Failed to connect: %v", err)}.Error()
	}
	defer client.Disconnect(context.TODO())
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		logger.Log{Message: fmt.Sprintf("Failed to connect to cluster: %v", err)}.Fatal()
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Log{Message: fmt.Sprintf("Failed to ping cluster: %v", err)}.Fatal()
	}
	logger.Log{Message: "Connected to DocumentDB!"}.Success()
	database := client.Database(databaseName)
	mongoDB = database
	fmt.Println("Successfully connected and pinged.")
}
