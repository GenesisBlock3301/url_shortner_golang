package db

import (
	"context"
	"fmt"
	"github.com/GenesisBlock3301/url_shortner_golang/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var (
	MongoDB *mongo.Database
	client  *mongo.Client
	err     error
)

const (
	connectTimeout = 30
)

func ConnectDB() {
	logger.Log{Message: fmt.Sprintf("Mongo url: %v", MongoUrl)}.Info()
	client, err = mongo.NewClient(options.Client().ApplyURI(MongoUrl))
	if err != nil {
		logger.Log{Message: fmt.Sprintf("Failed to connect: %v", err)}.Error()
	}
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
	MongoDB = client.Database(DatabaseName)
	fmt.Println("Database successfully connected and pinged.")
}
