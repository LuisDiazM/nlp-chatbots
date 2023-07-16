package controllers

import (
	"github.com/LuisDiazM/agent-manager/domain/usecases/trainingUsecase/entities"
	"github.com/LuisDiazM/agent-manager/infraestructure/app"

	"net/http"

	"github.com/gin-gonic/gin"
)

func TrainingModelInfo(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		trainingData := app.TrainingUsecase.GetModelById(id, ctx.Request.Context())
		if trainingData != nil {
			ctx.JSON(http.StatusOK, trainingData)
			return
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}

func InsertTrainingModelInfo(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var trainingInfo entities.TrainingInfo
		err := ctx.ShouldBindJSON(&trainingInfo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		} else {
			response := app.TrainingUsecase.InsertModel(trainingInfo, ctx.Request.Context())
			if response == nil {
				ctx.JSON(http.StatusNoContent, gin.H{})
				return
			} else {
				ctx.JSON(http.StatusCreated, gin.H{"model_id": response})
			}
		}
	}
}

func DeleteTrainingModelInfo(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		err := app.TrainingUsecase.DeleteTrainingIntentById(id, ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusNoContent, nil)
		}
	}
}

func UpdateTrainingModelInfoById(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, _ := ctx.Params.Get("id")
		var trainingInfo entities.TrainingInfo
		err := ctx.ShouldBindJSON(&trainingInfo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		_, err = app.TrainingUsecase.UpdateTrainingIntentById(id, trainingInfo, ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
		} else {
			ctx.JSON(http.StatusOK, gin.H{"model_id": id})
		}
	}
}
