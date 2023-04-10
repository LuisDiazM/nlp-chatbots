package routes

import (
	"http-models-server/infraestructure/server/controllers"

	"github.com/gin-gonic/gin"
)

func NewHealtRouter(group *gin.RouterGroup) {
	group.GET("/health", controllers.HealthController)
}
