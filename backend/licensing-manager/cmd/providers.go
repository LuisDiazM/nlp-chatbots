package cmd

import (
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/cmd/config"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/usecases"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/app"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/database"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/database/repositories"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/messaging"
	"github.com/google/wire"
)

var AppProvider = wire.NewSet(app.NewApplication)
var EnvironmentVariablesProvider = wire.NewSet(config.NewEnvironmentsSpecification)
var NatsProvider = wire.NewSet(messaging.NewNatsImplementation)
var DatabaseProvider = wire.NewSet(database.NewDatabaseImplementation)

var LicenseRepoProvider = wire.NewSet(repositories.NewLicenseRepository)
var LicenseUsecaseProvider = wire.NewSet(usecases.NewLicenseUseCase)
