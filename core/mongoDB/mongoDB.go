package mongoDB

import (
	"context"
	"time"

	"example.com/test/core/initializers"
	"example.com/test/core/logger"
	"example.com/test/core/modify"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func collection(name string) (*mongo.Collection, context.Context) {
	ctx := context.TODO()
	coll := initializers.MongoDB.Collection(name)

	return coll, ctx
}

func CID(id string) primitive.ObjectID {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error(err)
		return primitive.NilObjectID
	}

	return objectID
}

func Insert[T any](collectionName string, data T) (id string, err error) {
	dataValue := modify.Map(data, map[string]any{
		"CreatedAt": time.Now(),
		"UpdatedAt": time.Now(),
	})

	coll, ctx := collection(collectionName)

	result, err := coll.InsertOne(ctx, dataValue)

	if err != nil {
		logger.Error(err)
		return "", err
	}

	id = result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func Find[T any](collectionName string) ([]T, error) {
	coll, ctx := collection(collectionName)

	query, err := coll.Find(ctx, bson.D{})
	if err != nil {
		logger.Print(err.Error())
		return nil, err
	}
	defer query.Close(ctx)

	data := make([]T, 0)
	err = query.All(ctx, &data)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return data, nil
}

func FindOne[T any](collectionName string, filter map[string]any) (data T, exist bool, err error) {
	coll, ctx := collection(collectionName)

	result := coll.FindOne(ctx, filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return data, false, nil
		}

		logger.Error(result.Err())
		return data, false, result.Err()
	}

	if err := result.Decode(&data); err != nil {
		logger.Error(err)
		return data, false, err
	}

	return data, true, nil
}

func FindByID[T any](collectionName string, id string) (data T, exist bool, err error) {
	getId := CID(id)

	if getId == primitive.NilObjectID {
		return data, false, nil
	}

	return FindOne[T](collectionName, map[string]any{
		"_id": getId,
	})
}

func Exists[T any](collectionName string, filter map[string]any) (bool, error) {
	coll, ctx := collection(collectionName)

	result := coll.FindOne(ctx, filter)

	if result.Err() != nil {
		return false, result.Err()
	}

	return true, nil
}
