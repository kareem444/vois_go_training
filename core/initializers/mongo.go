package initializers

import (
	"context"
	"time"

	"example.com/test/core/env"
	"example.com/test/core/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database

func InitMongo() {
	MONGODB_URI := env.Get("MONGODB_URI")
	MONGODB_DB_NAME := env.Get("MONGODB_DB_NAME")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(MONGODB_URI).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		logger.Error(err)
	}

	errPing := client.Database(MONGODB_DB_NAME).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	logger.Error(errPing)

	MongoDB = client.Database(MONGODB_DB_NAME)
	logger.Info("successfully connected to MongoDB!")
}
