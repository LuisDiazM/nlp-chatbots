// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/cmd/config"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/usecases"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/app"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/database"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/database/repositories"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/messaging"
)

// Injectors from wire.go:

func CreateApp() *app.Application {
	env := config.NewEnvironmentsSpecification()
	natsImp := messaging.NewNatsImplementation(env)
	databaseImp := database.NewDatabaseImplementation(env)
	licenseRepositoryGateway := repositories.NewLicenseRepository(databaseImp)
	licenseUsecase := usecases.NewLicenseUseCase(licenseRepositoryGateway)
	application := app.NewApplication(env, natsImp, databaseImp, licenseUsecase)
	return application
}
