package routes

import (
	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/controllers"
	"github.com/gin-gonic/gin"
)

func NewLicensesRouter(group *gin.RouterGroup, app *app.Application) {

	group.GET("/license", controllers.GetLastLicenseByUserId(app))
	group.GET("/license/usage", controllers.GetLastLicenseUsage(app))

}
