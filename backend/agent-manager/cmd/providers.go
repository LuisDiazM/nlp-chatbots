package cmd

import (
	"github.com/LuisDiazM/agent-manager/cmd/config"
	trainingusecase "github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase"
	trainingRepository "github.com/LuisDiazM/agent-manager/infraestructure/database/repositories"

	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/database"
	"github.com/LuisDiazM/agent-manager/infraestructure/server"

	"github.com/google/wire"
)

var AppProvider = wire.NewSet(app.NewApplication)
var WebServerProvider = wire.NewSet(server.NewServer)
var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var DatabaseProvider = wire.NewSet(database.NewDatabaseImplementation)

var TrainingUsecaseProvider = wire.NewSet(trainingusecase.NewTrainingUsecase)
var TrainingRepositoryProvider = wire.NewSet(trainingRepository.NewTrainingRepository)