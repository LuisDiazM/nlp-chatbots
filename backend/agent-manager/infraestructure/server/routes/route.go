package routes

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/middlewares"
)

func SetUpRoutes(app *app.Application) {
	publicRoutes := app.WebServer.Group("")
	NewHealtRouter(publicRoutes)
	NewLoginRouter(publicRoutes, app)
	privateRoutes := app.WebServer.Group("")
	privateRoutes.Use(middlewares.JwtGoogle())
	NewTrainingRouter(privateRoutes, app)
}
