package cmd

import (
	"github.com/LuisDiazM/agent-manager/cmd/config"
	trainingusecase "github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase"
	userusecase "github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase"
	repository "github.com/LuisDiazM/agent-manager/infraestructure/database/repositories"
	"github.com/LuisDiazM/agent-manager/infraestructure/messaging"
	messagingLicensesRepo "github.com/LuisDiazM/agent-manager/infraestructure/messaging/repositories/userRepository"

	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/database"
	"github.com/LuisDiazM/agent-manager/infraestructure/server"

	"github.com/google/wire"
)

var AppProvider = wire.NewSet(app.NewApplication)
var WebServerProvider = wire.NewSet(server.NewServer)
var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var DatabaseProvider = wire.NewSet(database.NewDatabaseImplementation)
var MessagingProvider = wire.NewSet(messaging.NewNatsImplementation)

var TrainingUsecaseProvider = wire.NewSet(trainingusecase.NewTrainingUsecase)
var TrainingRepositoryProvider = wire.NewSet(repository.NewTrainingRepository)

var UserRepositoryProvider = wire.NewSet(repository.NewUserRepository)
var UserLicensesRepositoryProvider = wire.NewSet(messagingLicensesRepo.NewUserLicenseMessagingRepository)
var UserUsecaseProvider = wire.NewSet(userusecase.NewUserUsecase)
