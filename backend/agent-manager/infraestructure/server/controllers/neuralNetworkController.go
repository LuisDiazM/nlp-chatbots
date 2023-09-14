package controllers

import (
	"net/http"

	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/gin-gonic/gin"
)

func GetNNModelsByTrainingId(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		trainingId := ctx.Query("trainingId")
		models, err := app.NeuralNetworkUsecase.GetModelsByTrainingId(trainingId, ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		if models != nil {
			ctx.JSON(http.StatusOK, models)
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}
