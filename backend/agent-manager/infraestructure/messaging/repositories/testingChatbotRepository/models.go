package testingChatbotRepository

import "time"

const (
	queryTestingChatbot = "get.chatbot.response.by.model.id"
	timeout             = 1 * time.Minute
)

type TestingChatbotRequest struct {
	Sentence string `json:"sentence,omitempty"`
	ModelId  string `json:"model_id,omitempty"`
	UserId   string `json:"user_id,omitempty"`
}
