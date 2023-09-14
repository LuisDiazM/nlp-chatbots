package nnModelsUsecase

import (
	"context"

	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/repositories"
)

type NeuralNetworkModelUsecase struct {
	DatabaseRepo repositories.NnModelsRepository
}

func NewNeuralNetworkModelUsecase(databaseRepo repositories.NnModelsRepository) *NeuralNetworkModelUsecase {
	return &NeuralNetworkModelUsecase{DatabaseRepo: databaseRepo}
}

func (usecase *NeuralNetworkModelUsecase) GetModelsByTrainingId(trainingId string, ctx context.Context) (*[]entities.NNModel, error) {
	return usecase.DatabaseRepo.GetModelsByTrainingId(ctx, trainingId)
}
