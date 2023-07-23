package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/cmd"
)

func main() {
	app := cmd.CreateApp()
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	app.Start(ctx)

}
