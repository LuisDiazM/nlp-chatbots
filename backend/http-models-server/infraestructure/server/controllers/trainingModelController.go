package controllers

import (
	trainingusecase "http-models-server/domain/usecases/training-usecase"
	"http-models-server/domain/usecases/training-usecase/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TrainingController struct {
	TrainingModelUsecase trainingusecase.TrainingUsecase
}

func (controller *TrainingController) TrainingModelInfo(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	trainingData := controller.TrainingModelUsecase.GetModelById(id, ctx.Request.Context())
	if trainingData != nil {
		ctx.JSON(http.StatusOK, trainingData)
		return
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}

func (controller *TrainingController) InsertTrainingModelInfo(ctx *gin.Context) {
	var trainingInfo entities.TrainingInfo
	err := ctx.ShouldBindJSON(&trainingInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} else {
		response := controller.TrainingModelUsecase.InsertModel(trainingInfo, ctx.Request.Context())
		if response == nil {
			ctx.JSON(http.StatusNoContent, gin.H{})
			return
		} else {
			ctx.JSON(http.StatusCreated, response)
		}
	}
}
