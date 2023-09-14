package repositories

import (
	"context"

	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/repositories"
	"github.com/LuisDiazM/agent-manager/infraestructure/database"
	"go.mongodb.org/mongo-driver/bson"
)

type NeuralNetworkModelRepository struct {
	Database *database.DatabaseImp
}

func NewNeuralNetworkRepository(db *database.DatabaseImp) repositories.NnModelsRepository {
	return &NeuralNetworkModelRepository{Database: db}
}

func (repository *NeuralNetworkModelRepository) GetModelsByTrainingId(ctx context.Context, trainingId string) (*[]entities.NNModel, error) {
	collection := repository.Database.Collection(trainingDatabaseName, nnModelsCollection)
	filter := bson.M{"nluIntentId": trainingId}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var models []entities.NNModel
	err = cursor.All(ctx, &models)
	if err != nil {
		return nil, err
	}
	return &models, nil
}
