package main

import (
	"context"
	"github.com/LuisDiazM/agent-manager/cmd"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/routes"
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
