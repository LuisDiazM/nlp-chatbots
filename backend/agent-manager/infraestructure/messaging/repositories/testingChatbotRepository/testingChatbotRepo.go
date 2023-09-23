package testingChatbotRepository

import (
	"encoding/json"
	"log"

	"github.com/LuisDiazM/agent-manager/domain/usecases/nnModelsUsecase/repositories"
	"github.com/LuisDiazM/agent-manager/infraestructure/messaging"
)

type TestingNNModelsRepository struct {
	Nats *messaging.NatsImp
}

func NewTestingNNModelsRepository(nats *messaging.NatsImp) repositories.NnModelsMessagingRepository {
	return &TestingNNModelsRepository{Nats: nats}
}

func (repository *TestingNNModelsRepository) TestingChatbot(content string, modelId string, userId string) *string {
	var testingChatbotRequest TestingChatbotRequest = TestingChatbotRequest{Sentence: content, ModelId: modelId, UserId: userId}
	data, err := json.Marshal(testingChatbotRequest)
	if err != nil {
		log.Println(err)
		return nil
	}
	msg, err := repository.Nats.Conn.Request(queryTestingChatbot, data, timeout)
	if err != nil {
		log.Println(err)
		return nil
	}
	chatResponse := string(msg.Data)
	return &chatResponse
}
