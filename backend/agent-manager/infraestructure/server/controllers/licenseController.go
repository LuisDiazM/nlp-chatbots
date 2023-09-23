package controllers

import (
	"net/http"

	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/gin-gonic/gin"
)

func GetLastLicenseByUserId(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Query("userId")
		response := app.UserUsecase.GetLastLicenseByUserId(userId)
		if response != nil {
			ctx.JSON(http.StatusOK, response)
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}

func GetLastLicenseUsage(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		licenseId := ctx.Query("licenseId")
		response := app.UserUsecase.GetLastLicenseUsage(licenseId)
		if response != nil {
			ctx.JSON(http.StatusOK, response)
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}
