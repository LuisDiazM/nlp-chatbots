package routes

import (
	"http-models-server/infraestructure/database"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(server *gin.Engine, databaseGateway *database.DatabaseImp) {
	publicRoutes := server.Group("")
	NewHealtRouter(publicRoutes)
	NewTrainingRouter(publicRoutes, *databaseGateway)
	NewModelRouter(publicRoutes, *databaseGateway)
}
