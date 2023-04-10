package main

import (
	"http-models-server/cmd"
)

func main() {
	app := cmd.CreateApp()

	app.Start()

}
