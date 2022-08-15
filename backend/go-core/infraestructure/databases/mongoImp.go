package databases

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/LuisDiazM/goCore/domain"
	"github.com/LuisDiazM/goCore/domain/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseChannel  = "channels"
	databaseModels   = "models"
	databaseTraining = "trainings"

	collectionChannelData  = "channelInfo"
	collectionTrainingData = "trainingInfo"
)

type DatabaseGatewayImp struct {
	client *mongo.Client
	ctx    context.Context
}

func NewDatabaseGatewayImp() domain.DatabaseGateway {
	return &DatabaseGatewayImp{}
}

func (database *DatabaseGatewayImp) Setup() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	opt := options.Client()
	opt.SetMaxPoolSize(2)
	opt.ApplyURI(os.Getenv("MONGO_URL"))
	database.client, err = mongo.Connect(ctx, opt)
	if err != nil {
		log.Fatal(err)
	}
}

func (database *DatabaseGatewayImp) Shutdown() {
	database.client.Disconnect(database.ctx)
}

func (database DatabaseGatewayImp) GetChannelsById(id string) (*models.ChannelInfo, error) {
	collection := database.client.Database(databaseChannel).Collection(collectionChannelData)
	ctx := context.TODO()
	var channelData models.ChannelInfo
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectId}
	cursor := collection.FindOne(ctx, filter)
	cursor.Decode(&channelData)
	return &channelData, nil
}

func (database DatabaseGatewayImp) InsertTrainingData(trainingInfo models.TrainingInfo) *interface{} {
	collection := database.client.Database(databaseTraining).Collection(collectionTrainingData)
	ctx := context.TODO()
	result, err := collection.InsertOne(ctx, trainingInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &result.InsertedID
}
