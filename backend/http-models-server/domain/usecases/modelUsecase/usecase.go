package modelusecase

import (
	"context"
	"http-models-server/domain/usecases/modelUsecase/entities"
	"http-models-server/domain/usecases/modelUsecase/repositories"
)

type ModelsUsecase struct {
	Repository repositories.ModelsRepository
}

func NewModelUsecase(repository repositories.ModelsRepository) *ModelsUsecase {
	return &ModelsUsecase{
		Repository: repository,
	}
}

func (usecase *ModelsUsecase) GetModelById(id string, ctx context.Context) *entities.Model {
	return usecase.Repository.GetModelNLP(id, ctx)
}
