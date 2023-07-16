package routes

import (
	"http-models-server/infraestructure/app"
)

func SetUpRoutes(app *app.Application) {
	publicRoutes := app.WebServer.Group("")
	NewHealtRouter(publicRoutes)
	NewTrainingRouter(publicRoutes, app)
}
