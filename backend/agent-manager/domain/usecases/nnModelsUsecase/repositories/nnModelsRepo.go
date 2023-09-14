package repositories

import (
	"context"

	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/entities"
)

type NnModelsRepository interface {
	GetModelsByTrainingId(ctx context.Context, trainingId string) (*[]entities.NNModel, error)
}
