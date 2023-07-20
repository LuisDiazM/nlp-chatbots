package repositories

import (
	"context"
	"log"

	"github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase/repositories"

	"github.com/LuisDiazM/agent-manager/infraestructure/database"
)

type TrainingRepository struct {
	Database *database.DatabaseImp
}

func NewTrainingRepository(db *database.DatabaseImp) repositories.TrainingRepository {
	return &TrainingRepository{
		Database: db,
	}
}

func (repository *TrainingRepository) GetTrainingModelById(id string, ctx context.Context) *entities.TrainingInfo {
	collection := repository.Database.Collection(trainingDatabaseName, trainingCollectionName)
	var trainingInfo entities.TrainingInfo
	data := repository.Database.FindOne(collection, &ctx, id)
	err := data.Decode(&trainingInfo)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &trainingInfo
}

func (repository *TrainingRepository) InsertTrainingModel(data entities.TrainingInfo, ctx context.Context) *interface{} {
	collection := repository.Database.Collection(trainingDatabaseName, trainingCollectionName)
	result := repository.Database.InsertOne(collection, &ctx, data)
	if result != nil {
		return &result.InsertedID
	} else {
		return nil
	}
}

func (repository *TrainingRepository) DeleteTrainingModel(id string, ctx context.Context) error {
	collection := repository.Database.Collection(trainingDatabaseName, trainingCollectionName)
	return repository.Database.DeleteOne(collection, &ctx, id)
}

func (repository *TrainingRepository) UpdateTrainingModel(id string, data entities.TrainingInfo, ctx context.Context) (*interface{}, error) {
	collection := repository.Database.Collection(trainingDatabaseName, trainingCollectionName)
	return repository.Database.UpdateOneById(collection, &ctx, id, data)

}