package trainingusecase

import (
	"context"
	"http-models-server/domain/usecases/trainingUsecase/entities"
	"http-models-server/domain/usecases/trainingUsecase/repositories"
)

type TrainingUsecase struct {
	TrainingRepository repositories.TrainingRepository
}

func NewTrainingUsecase(trainingRepository repositories.TrainingRepository) *TrainingUsecase {
	return &TrainingUsecase{TrainingRepository: trainingRepository}
}

func (usecase *TrainingUsecase) GetModelById(id string, ctx context.Context) *entities.TrainingInfo {
	return usecase.TrainingRepository.GetTrainingModelById(id, ctx)
}

func (usecase *TrainingUsecase) InsertModel(data entities.TrainingInfo, ctx context.Context) *interface{} {
	return usecase.TrainingRepository.InsertTrainingModel(data, ctx)
}

func (usecase *TrainingUsecase) DeleteTrainingIntentById(id string, ctx context.Context) error {
	return usecase.TrainingRepository.DeleteTrainingModel(id, ctx)
}

func (usecase *TrainingUsecase) UpdateTrainingIntentById(id string, data entities.TrainingInfo, ctx context.Context) (*interface{}, error) {
	return usecase.TrainingRepository.UpdateTrainingModel(id, data, ctx)
}