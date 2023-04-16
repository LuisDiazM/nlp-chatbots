package controllers

import (
	"net/http"

	modelusecase "http-models-server/domain/usecases/modelUsecase"

	"github.com/gin-gonic/gin"
)

type ModelController struct {
	ModelUsecase modelusecase.ModelsUsecase
}

func (controller *ModelController) GetModelNLPController(ctx *gin.Context) {
	id := ctx.Param("id")
	model := controller.ModelUsecase.GetModelById(id, ctx.Request.Context())
	if model != nil {
		ctx.JSON(http.StatusOK, model)
	} else {
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
}
