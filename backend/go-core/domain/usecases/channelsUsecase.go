package usecases

import (
	"github.com/LuisDiazM/goCore/domain"
	"github.com/LuisDiazM/goCore/domain/models"
)

type ChannelsUsecase struct {
	databaseGateway domain.DatabaseGateway
}

func NewChannelUsecase(databaseUsecase domain.DatabaseGateway) ChannelsUsecase {
	return ChannelsUsecase{
		databaseGateway: databaseUsecase,
	}
}

func (channel ChannelsUsecase) GetChannelById(id string) (*models.ChannelInfo, error) {
	channelInfo, err := channel.databaseGateway.GetChannelsById(id)
	return channelInfo, err
}
