package repositories

type TrainingMessagingRepository interface {
	CreateNNModel(userId string, nluIntentId string) error
	DeleteNNModelsByNluIntentId(nluIntentId string) error
	DeleteNNModelsByUserId(userId string) error
}
