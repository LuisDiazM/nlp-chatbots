package app

import (
	"log"
	"net/http"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
)

type Application struct {
	Router *mux.Router
}

func NewApplication() *Application {
	return &Application{}
}

func (app Application) Start() {
	withGz := gziphandler.GzipHandler(app.Router)
	http.Handle("/", withGz)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
