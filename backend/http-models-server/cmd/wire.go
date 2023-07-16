//go:build wireinject
// +build wireinject

package cmd

import (
	"http-models-server/infraestructure/app"

	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(AppProvider,
		WebServerProvider,
		EnvironmentVariablesProvider,
		DatabaseProvider,
		TrainingUsecaseProvider,
		TrainingRepositoryProvider,
	)
	return new(app.Application)
}
