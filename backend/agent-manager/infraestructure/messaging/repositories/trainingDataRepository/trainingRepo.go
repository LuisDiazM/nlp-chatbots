package trainingDataRepository

import (
	"encoding/json"
	"fmt"

	"github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase/repositories"
	"github.com/LuisDiazM/agent-manager/infraestructure/messaging"
)

type TrainingNNModelsRepository struct {
	Nats *messaging.NatsImp
}

func NewTrainingNNModelsRepository(nats *messaging.NatsImp) repositories.TrainingMessagingRepository {
	return &TrainingNNModelsRepository{Nats: nats}
}

func (repository *TrainingNNModelsRepository) CreateNNModel(userId string, nluIntentId string) error {
	requestData := TrainModelRequest{UserId: userId, Id: nluIntentId}
	data, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = repository.Nats.Conn.Publish(eventCreateNeuralNetworkModel, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

func (repository *TrainingNNModelsRepository) DeleteNNModelsByNluIntentId(nluIntentId string) error {
	requestData := ModelIdRequest{ModelId: nluIntentId}
	data, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = repository.Nats.Conn.Publish(eventDeleteNeuralNetworkModel, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (repository *TrainingNNModelsRepository) DeleteNNModelsByUserId(userId string) error {
	requestData := UserIdRequest{UserId: userId}
	data, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = repository.Nats.Conn.Publish(eventDeleteNeuralNetworkModelByUserId, data)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}
