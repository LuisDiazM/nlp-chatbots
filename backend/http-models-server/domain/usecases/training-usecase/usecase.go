package trainingusecase

import (
	"context"
	"http-models-server/domain/usecases/training-usecase/entities"
	"http-models-server/domain/usecases/training-usecase/repositories"
)

type TrainingUsecase struct {
	TrainingRepository repositories.TrainingRepository
}

func NewTrainingUsecase(trainingRepository repositories.TrainingRepository) *TrainingUsecase {
	return &TrainingUsecase{TrainingRepository: trainingRepository}
}

func (usecase *TrainingUsecase) GetModelById(id string, ctx context.Context) *entities.TrainingInfo {
	return usecase.TrainingRepository.GetModelById(id, ctx)

}
