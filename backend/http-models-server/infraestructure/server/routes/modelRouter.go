package routes

import (
	modelusecase "http-models-server/domain/usecases/modelUsecase"
	"http-models-server/infraestructure/database"
	"http-models-server/infraestructure/database/repositories"
	"http-models-server/infraestructure/server/controllers"

	"github.com/gin-gonic/gin"
)

func NewModelRouter(group *gin.RouterGroup, databaseGateway database.DatabaseImp) {
	repository := repositories.NewModelRepository(databaseGateway)
	controller := &controllers.ModelController{
		ModelUsecase: *modelusecase.NewModelUsecase(repository),
	}
	group.GET("/models/:id", controller.GetModelNLPController)
}
