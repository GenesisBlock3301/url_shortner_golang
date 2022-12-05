package helpers

import (
	"context"
	"github.com/GenesisBlock3301/url_shortner_golang/config/db"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

//
func GetDB(collectionName string) *mongo.Collection {
	database := db.MongoDB
	collection := database.Collection(collectionName)
	return collection
}

func GetContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5000*time.Second)
	return ctx, cancel
}
