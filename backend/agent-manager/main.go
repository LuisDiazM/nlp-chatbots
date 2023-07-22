package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/agent-manager/cmd"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/middlewares"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/routes"
)

func main() {
	app := cmd.CreateApp()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	middlewares.EnableCors(app)
	routes.SetUpRoutes(app)
	err := app.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

}
