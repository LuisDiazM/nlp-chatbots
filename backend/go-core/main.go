package main

import (
	"github.com/LuisDiazM/goCore/app"
	"github.com/gorilla/mux"
)

func main() {
	app := app.CreateApp()
	app.Router = mux.NewRouter()
	app.Setup()
	app.Start()
}
