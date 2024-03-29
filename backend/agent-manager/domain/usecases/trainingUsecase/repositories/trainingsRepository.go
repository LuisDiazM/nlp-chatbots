package repositories

import (
	"context"

	"github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase/entities"
)

type TrainingRepository interface {
	GetTrainingModelById(id string, ctx context.Context) *entities.TrainingInfo
	InsertTrainingModel(data entities.TrainingInfo, ctx context.Context) string
	DeleteTrainingModel(id string, ctx context.Context) error
	UpdateTrainingModel(id string, data entities.TrainingInfo, ctx context.Context) (*interface{}, error)
	GetModelsByUserId(ctx context.Context, userId string) (*[]entities.TrainingInfo, error)
}
