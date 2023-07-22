//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/app"

	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(AppProvider,
		WebServerProvider,
		EnvironmentVariablesProvider,
		DatabaseProvider,
		TrainingUsecaseProvider,
		TrainingRepositoryProvider,
		UserRepositoryProvider,
		UserUsecaseProvider,
	)
	return new(app.Application)
}
