package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/LuisDiazM/goCore/domain/models"
)

// @Summary Index
// @Description This endpoint does not do nothing
// @Tags Index
// @Success 200 {object} object
// @Failure 401 {object} string
// @Accept  json
// @Produce  json
// @Router / [get]
func (app Application) Index(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusOK)
}

// @Summary
// @Description This endpoint does not do nothing
// @Tags Training
// @Success 200 {object} models.ObjectCreated
// @Failure 401 {object} string
// @Accept  json
// @Produce  json
// @Param default body models.TrainingInfo true "training data"
// @Router /training-data [post]
// @Security ApiKeyAuth
func (app Application) SaveTrainingData(response http.ResponseWriter, request *http.Request) {
	requestBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	var trainingData models.TrainingInfo
	err = json.Unmarshal(requestBytes, &trainingData)
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}
	docID := app.trainingUsecase.SaveTrainingData(trainingData)
	var objectData models.ObjectCreated = models.ObjectCreated{Id: docID}
	data, err := json.Marshal(objectData)
	response.Header().Add("Content-Type", "application/json")
	response.Write(data)
}
