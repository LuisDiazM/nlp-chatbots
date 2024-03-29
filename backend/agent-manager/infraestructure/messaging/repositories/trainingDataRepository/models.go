package trainingDataRepository

type TrainModelRequest struct {
	Id     string `json:"id,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

type ModelIdRequest struct {
	ModelId string `json:"training_id,omitempty"`
}

type UserIdRequest struct {
	UserId string `json:"user_id,omitempty"`
}
