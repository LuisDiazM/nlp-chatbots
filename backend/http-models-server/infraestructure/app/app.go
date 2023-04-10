package app

import (
	"fmt"
	"http-models-server/cmd/config"
	"http-models-server/infraestructure/database"
	"http-models-server/infraestructure/server/routes"

	"github.com/gin-gonic/gin"
)

type Application struct {
	WebServer *gin.Engine
	Env       *config.Env
	Database  *database.DatabaseImp
}

func NewApplication(webServer *gin.Engine, configVars *config.Env, database *database.DatabaseImp) *Application {
	return &Application{WebServer: webServer,
		Env:      configVars,
		Database: database,
	}
}

func (app *Application) Start() {
	server := app.WebServer
	app.Database.Setup(*app.Env)
	routes.SetUpRoutes(server, app.Database)
	server.Run(fmt.Sprintf(":%d", app.Env.Port))
}
