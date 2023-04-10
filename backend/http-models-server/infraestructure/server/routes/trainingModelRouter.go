package routes

import (
	trainingusecase "http-models-server/domain/usecases/training-usecase"
	"http-models-server/infraestructure/database"
	"http-models-server/infraestructure/database/repositories"

	"http-models-server/infraestructure/server/controllers"

	"github.com/gin-gonic/gin"
)

func NewTrainingRouter(group *gin.RouterGroup, databaseGateway database.DatabaseImp) {
	repository := repositories.NewTrainingRepository(databaseGateway)
	controller := &controllers.TrainingController{
		TrainingModelUsecase: *trainingusecase.NewTrainingUsecase(repository),
	}
	group.GET("/training-model/:id", controller.ModelInfo)
}
