package cmd

import (
	"http-models-server/cmd/config"
	"http-models-server/infraestructure/app"
	"http-models-server/infraestructure/database"
	"http-models-server/infraestructure/server"

	"github.com/google/wire"
)

var AppProvider = wire.NewSet(app.NewApplication)
var WebServerProvider = wire.NewSet(server.NewServer)
var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var DatabaseProvider = wire.NewSet(database.NewDatabaseImplementation)

// var TrainingUsecaseProvider = wire.NewSet(trainingusecase.NewTrainingUsecase)
