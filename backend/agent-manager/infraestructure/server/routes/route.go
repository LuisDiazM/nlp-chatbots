package routes

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/app"
)

func SetUpRoutes(app *app.Application) {
	publicRoutes := app.WebServer.Group("")
	NewHealtRouter(publicRoutes)
	NewTrainingRouter(publicRoutes, app)
}
