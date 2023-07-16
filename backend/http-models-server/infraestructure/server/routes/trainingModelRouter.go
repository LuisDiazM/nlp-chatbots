package routes

import (
	"http-models-server/infraestructure/app"
	"http-models-server/infraestructure/server/controllers"

	"github.com/gin-gonic/gin"
)

func NewTrainingRouter(group *gin.RouterGroup, app *app.Application) {

	group.GET("/training-model/:id", controllers.TrainingModelInfo(app))
	group.POST("/training-model", controllers.InsertTrainingModelInfo(app))
	group.DELETE("/training-model/:id", controllers.DeleteTrainingModelInfo(app))
	group.PUT("/training-model/:id", controllers.UpdateTrainingModelInfoById(app))

}
