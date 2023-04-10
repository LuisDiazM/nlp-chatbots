package controllers

import (
	trainingusecase "http-models-server/domain/usecases/training-usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainingController struct {
	TrainingModelUsecase trainingusecase.TrainingUsecase
}

func (controller *TrainingController) ModelInfo(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	trainingData := controller.TrainingModelUsecase.GetModelById(id, ctx.Request.Context())
	if trainingData != nil {
		ctx.JSON(http.StatusOK, trainingData)
		return
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
