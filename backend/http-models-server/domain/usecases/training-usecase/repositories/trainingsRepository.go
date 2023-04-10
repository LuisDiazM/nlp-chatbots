package repositories

import (
	"context"
	"http-models-server/domain/usecases/training-usecase/entities"
)

type TrainingRepository interface {
	GetModelById(id string, ctx context.Context) *entities.TrainingInfo
}
