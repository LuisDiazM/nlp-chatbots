package routes

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/controllers"

	"github.com/gin-gonic/gin"
)

func NewTrainingRouter(group *gin.RouterGroup, app *app.Application) {

	group.GET("/training-model/:id", controllers.TrainingModelInfo(app))
	group.POST("/training-model", controllers.InsertTrainingModelInfo(app))
	group.DELETE("/training-model/:id", controllers.DeleteTrainingModelInfo(app))
	group.PUT("/training-model/:id", controllers.UpdateTrainingModelInfoById(app))
	group.GET("training-model", controllers.GetModelsByUserId(app))

}
