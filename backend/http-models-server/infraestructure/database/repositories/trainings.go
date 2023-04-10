package repositories

import (
	"context"
	"encoding/json"
	"fmt"
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
	data := repository.Database.FindOne(collection, &ctx, id)
	fmt.Println(*data)
	var trainingInfo entities.TrainingInfo
	bytesData, err := json.Marshal(*data)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(bytesData, &trainingInfo)
	if err != nil {
		log.Println("Repository GetModelById", err)
		return nil
	}

	return &trainingInfo
}
