package domain

import "github.com/LuisDiazM/goCore/domain/models"

type DatabaseGateway interface {
	Setup()
	Shutdown()
	GetChannelsById(id string) (*models.ChannelInfo, error)
	InsertTrainingData(trainingInfo models.TrainingInfo) *interface{}
}
