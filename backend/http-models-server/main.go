package main

import (
	"context"
	"http-models-server/cmd"
	"http-models-server/infraestructure/server/routes"
	"log"
	"os/signal"
	"syscall"
)

func main() {
	app := cmd.CreateApp()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	routes.SetUpRoutes(app)
	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
