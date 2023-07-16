package routes

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/server/controllers"

	"github.com/gin-gonic/gin"
)

func NewHealtRouter(group *gin.RouterGroup) {
	group.GET("/health", controllers.HealthController)
}
