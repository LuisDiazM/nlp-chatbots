//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/app"
	"github.com/google/wire"
)

func CreateApp() *app.Application {
	wire.Build(AppProvider,
		EnvironmentVariablesProvider,
		NatsProvider,
		DatabaseProvider,
		LicenseRepoProvider,
		LicenseUsecaseProvider,
	)
	return new(app.Application)
}
