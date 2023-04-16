package repositories

import (
	"context"
	"http-models-server/domain/usecases/training-usecase/entities"
	"http-models-server/domain/usecases/training-usecase/repositories"
	"log"

	"http-models-server/infraestructure/database"
)

type TrainingRepository struct {
	Database database.DatabaseImp
}

func NewTrainingRepository(db database.DatabaseImp) repositories.TrainingRepository {
	return &TrainingRepository{
		Database: db,
	}
}

func (repository *TrainingRepository) GetTrainingModelById(id string, ctx context.Context) *entities.TrainingInfo {
	collection := repository.Database.Collection(trainingDatabaseName, trainingCollectionName)
	var trainingInfo *entities.TrainingInfo
	data := repository.Database.FindOne(collection, &ctx, id)
	err := data.Decode(&trainingInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	return trainingInfo
}

func (repository *TrainingRepository) InsertTrainingModel(data entities.TrainingInfo, ctx context.Context) *interface{} {
	collection := repository.Database.Collection(trainingCollectionName, trainingCollectionName)
	result := repository.Database.InsertOne(collection, &ctx, data)
	if result != nil {
		return &result.InsertedID
	} else {
		return nil
	}
}
