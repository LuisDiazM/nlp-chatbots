package routes

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/controllers"
	"github.com/gin-gonic/gin"
)

func NewNeuralNetworkRouter(group *gin.RouterGroup, app *app.Application) {

	group.GET("/nn-models", controllers.GetNNModelsByTrainingId(app))

}
