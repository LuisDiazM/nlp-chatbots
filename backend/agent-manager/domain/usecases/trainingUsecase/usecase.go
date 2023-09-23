package trainingusecase

import (
	"context"
	"fmt"

	"github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase/entities"
	"github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase/repositories"
)

type TrainingUsecase struct {
	TrainingRepository          repositories.TrainingRepository
	TrainingMessagingRepository repositories.TrainingMessagingRepository
}

func NewTrainingUsecase(trainingRepository repositories.TrainingRepository, trainingMessagingRepo repositories.TrainingMessagingRepository) *TrainingUsecase {
	return &TrainingUsecase{TrainingRepository: trainingRepository, TrainingMessagingRepository: trainingMessagingRepo}
}

func (usecase *TrainingUsecase) GetModelById(id string, ctx context.Context) *entities.TrainingInfo {
	return usecase.TrainingRepository.GetTrainingModelById(id, ctx)
}

func (usecase *TrainingUsecase) InsertModel(data entities.TrainingInfo, ctx context.Context) string {
	response := usecase.TrainingRepository.InsertTrainingModel(data, ctx)
	if response != "" {
		usecase.TrainingMessagingRepository.CreateNNModel(data.UserId, response)
	}
	return response
}

func (usecase *TrainingUsecase) DeleteTrainingIntentById(id string, ctx context.Context) error {
	errServiceTraining := usecase.TrainingMessagingRepository.DeleteNNModelsByNluIntentId(id)
	if errServiceTraining != nil {
		fmt.Println(errServiceTraining)
	}
	err := usecase.TrainingRepository.DeleteTrainingModel(id, ctx)
	if err != nil {
		fmt.Println(err)
	}
	if err == nil || errServiceTraining == nil {
		return nil
	}
	return fmt.Errorf(`model %s could not be deleted`, id)

}

func (usecase *TrainingUsecase) UpdateTrainingIntentById(id string, data entities.TrainingInfo, ctx context.Context) (*interface{}, error) {
	response, err := usecase.TrainingRepository.UpdateTrainingModel(id, data, ctx)
	if err != nil {
		return nil, err
	}
	usecase.TrainingMessagingRepository.CreateNNModel(data.UserId, id)
	return response, err
}

func (usecase *TrainingUsecase) GetModelsByUserId(userId string, ctx context.Context) (*[]entities.TrainingInfo, error) {
	return usecase.TrainingRepository.GetModelsByUserId(ctx, userId)
}

func (usecase *TrainingUsecase) DeleteNNModelsByUserId(userId string, ctx context.Context) error {
	models, _ := usecase.TrainingRepository.GetModelsByUserId(ctx, userId)
	for _, model := range *models {
		usecase.TrainingRepository.DeleteTrainingModel(model.Id, ctx)
	}
	return usecase.TrainingMessagingRepository.DeleteNNModelsByUserId(userId)
}
