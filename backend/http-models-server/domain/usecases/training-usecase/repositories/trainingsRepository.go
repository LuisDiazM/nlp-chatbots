package repositories

import (
	"context"
	"http-models-server/domain/usecases/training-usecase/entities"
)

type TrainingRepository interface {
	GetTrainingModelById(id string, ctx context.Context) *entities.TrainingInfo
	InsertTrainingModel(data entities.TrainingInfo, ctx context.Context) *interface{}
}
