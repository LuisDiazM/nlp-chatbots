package routes

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/controllers"
	"github.com/gin-gonic/gin"
)

func NewLoginRouter(group *gin.RouterGroup, app *app.Application) {

	group.GET("/login", controllers.LoginController(app))
	group.POST("/register", controllers.RegisterController(app))

}
