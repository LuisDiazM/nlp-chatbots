package nnModelsUsecase

import (
	"context"

	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/repositories"
)

type NeuralNetworkModelUsecase struct {
	DatabaseRepo  repositories.NnModelsRepository
	MessagingRepo repositories.NnModelsMessagingRepository
}

func NewNeuralNetworkModelUsecase(databaseRepo repositories.NnModelsRepository, messagingRepo repositories.NnModelsMessagingRepository) *NeuralNetworkModelUsecase {
	return &NeuralNetworkModelUsecase{DatabaseRepo: databaseRepo, MessagingRepo: messagingRepo}
}

func (usecase *NeuralNetworkModelUsecase) GetModelsByTrainingId(trainingId string, ctx context.Context) (*[]entities.NNModel, error) {
	return usecase.DatabaseRepo.GetModelsByTrainingId(ctx, trainingId)
}

func (usecase *NeuralNetworkModelUsecase) GetChatbotResponsesByModelId(content string, modelId string, userId string) entities.ChatbotResponse {
	response := usecase.MessagingRepo.TestingChatbot(content, modelId, userId)
	return entities.ChatbotResponse{ChatReponse: *response}
}
