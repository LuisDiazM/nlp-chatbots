//go:build wireinject
// +build wireinject

package app

import "github.com/google/wire"

func CreateApp() *Application {
	wire.Build(appProvider, databaseProvider, channelUsecaseProvider, trainingUsecaseProvider)
	return new(Application)
}
