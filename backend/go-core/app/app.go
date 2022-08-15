package app

import (
	"log"
	"net/http"

	"github.com/LuisDiazM/goCore/domain"
	"github.com/LuisDiazM/goCore/domain/usecases"
	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
)

type Application struct {
	Router          *mux.Router
	database        domain.DatabaseGateway
	channelUsecase  usecases.ChannelsUsecase
	trainingUsecase usecases.TrainingUseCase
}

func NewApplication(database domain.DatabaseGateway, channelUsecase usecases.ChannelsUsecase, trainingUsecase usecases.TrainingUseCase) *Application {
	return &Application{
		database:        database,
		channelUsecase:  channelUsecase,
		trainingUsecase: trainingUsecase,
	}
}

func (app Application) Start() {
	withGz := gziphandler.GzipHandler(app.Router)
	http.Handle("/", withGz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
