package app

import (
	"github.com/LuisDiazM/goCore/domain/usecases"
	"github.com/LuisDiazM/goCore/infraestructure/databases"
	"github.com/google/wire"
)

var appProvider = wire.NewSet(NewApplication)
var channelUsecaseProvider = wire.NewSet(usecases.NewChannelUsecase)
var trainingUsecaseProvider = wire.NewSet(usecases.NewTrainingUsecase)
var databaseProvider = wire.NewSet(databases.NewDatabaseGatewayImp)
