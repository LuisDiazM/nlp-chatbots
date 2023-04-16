package repositories

import (
	"context"
	"http-models-server/domain/usecases/training-usecase/entities"
	"http-models-server/domain/usecases/training-usecase/repositories"
	"log"

	"http-models-server/infraestructure/database"
)

const (
	databaseName   = "trainings"
	collectionName = "trainingInfo"
)

type TrainingRepository struct {
	Database database.DatabaseImp
}

func NewTrainingRepository(db database.DatabaseImp) repositories.TrainingRepository {
	return &TrainingRepository{
		Database: db,
	}
}

func (repository *TrainingRepository) GetModelById(id string, ctx context.Context) *entities.TrainingInfo {
	collection := repository.Database.Collection(databaseName, collectionName)
	var trainingInfo *entities.TrainingInfo
	data := repository.Database.FindOne(collection, &ctx, id)
	err := data.Decode(&trainingInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	return trainingInfo
}
