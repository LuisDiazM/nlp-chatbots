package database

import (
	"context"
	"log"
	"time"

	"http-models-server/cmd/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseImp struct {
	Client *mongo.Client
}

func NewDatabaseImplementation(environment *config.Env) *DatabaseImp {
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	opt := options.Client()
	opt.SetMaxPoolSize(environment.MONGO_POOL_SIZE)
	opt.ApplyURI(environment.MONGO_URL)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
	return &DatabaseImp{Client: client}
}

func (database *DatabaseImp) Collection(databaseName string, colName string) *mongo.Collection {
	return database.Client.Database(databaseName).Collection(colName)
}

func (database *DatabaseImp) FindOne(collection *mongo.Collection, ctx *context.Context, id string) *mongo.SingleResult {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("mongo -> FindOne", err)
		return nil
	}
	filter := bson.M{"_id": objectId}
	cursor := collection.FindOne(*ctx, filter)
	return cursor
}

func (database *DatabaseImp) InsertOne(collection *mongo.Collection, ctx *context.Context, data any) *mongo.InsertOneResult {
	result, err := collection.InsertOne(*ctx, data)
	if err != nil {
		log.Println("mongo.go -> InsertOne", err)
		return nil
	}
	return result
}

func (database *DatabaseImp) DeleteOne(collection *mongo.Collection, ctx *context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("mongo -> DeleteOne", err)
		return nil
	}
	filter := bson.M{"_id": objectId}

	_, err = collection.DeleteOne(*ctx, filter)
	if err != nil {
		log.Println("mongo -> DeleteOne", err)
		return nil
	}
	return nil
}

func (database *DatabaseImp) UpdateOneById(collection *mongo.Collection, ctx *context.Context, id string, data any) (*interface{}, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("mongo -> UpdateOneById", err)
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": data,
	}
	result, err := collection.UpdateOne(*ctx, filter, update)
	return &result.UpsertedID, err
}
