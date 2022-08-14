//go:build wireinject
// +build wireinject

package app

import "github.com/google/wire"

func CreateApp() *Application {
	wire.Build(appProvider)
	return new(Application)
}
