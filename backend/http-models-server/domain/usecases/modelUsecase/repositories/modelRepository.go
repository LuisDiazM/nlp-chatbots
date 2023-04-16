package repositories

import (
	"context"
	"http-models-server/domain/usecases/modelUsecase/entities"
)

type ModelsRepository interface {
	GetModelNLP(id string, ctx context.Context) *entities.Model
}
