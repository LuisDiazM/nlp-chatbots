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

func NewDatabaseImplementation() *DatabaseImp {
	return &DatabaseImp{}
}

func (database *DatabaseImp) Setup(environment config.Env) {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	opt := options.Client()
	opt.SetMaxPoolSize(uint64(environment.MONGO_POOL_SIZE))
	opt.ApplyURI(environment.MONGO_URL)
	database.Client, err = mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
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
