package usecases

import (
	"github.com/LuisDiazM/goCore/domain"
	"github.com/LuisDiazM/goCore/domain/models"
)

type TrainingUseCase struct {
	databaseGateway domain.DatabaseGateway
}

func NewTrainingUsecase(databaseUsecase domain.DatabaseGateway) TrainingUseCase {
	return TrainingUseCase{
		databaseGateway: databaseUsecase,
	}
}

func (training TrainingUseCase) SaveTrainingData(trainingInfo models.TrainingInfo) *interface{} {
	docId := training.databaseGateway.InsertTrainingData(trainingInfo)
	return docId
}
