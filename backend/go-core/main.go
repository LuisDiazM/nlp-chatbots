package main

import (
	"github.com/LuisDiazM/goCore/app"
	"github.com/gorilla/mux"
)

// @title Backend
// @version 1.0.0
// @description This microservice connects the front-end and is the core of the project
// @BasePath /api/v1
func main() {
	app := app.CreateApp()
	app.Router = mux.NewRouter()
	app.Setup()
	app.Start()
}
