package repositories

import (
	"context"
	"http-models-server/domain/usecases/trainingUsecase/entities"
)

type TrainingRepository interface {
	GetTrainingModelById(id string, ctx context.Context) *entities.TrainingInfo
	InsertTrainingModel(data entities.TrainingInfo, ctx context.Context) *interface{}
	DeleteTrainingModel(id string, ctx context.Context) error
	UpdateTrainingModel(id string, data entities.TrainingInfo, ctx context.Context) (*interface{}, error)
}
