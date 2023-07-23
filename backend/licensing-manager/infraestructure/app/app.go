package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/cmd/config"
	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/infraestructure/messaging"
)

type Application struct {
	Env  *config.Env
	Nats *messaging.NatsImp
}

func NewApplication(envs *config.Env, natsImp *messaging.NatsImp) *Application {
	return &Application{
		Env:  envs,
		Nats: natsImp,
	}
}

func (app *Application) Start(ctx context.Context) {
	defer app.Nats.CloseConnection()
	log.Println("Start app licensing-manager")
	app.RunSubscribers()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Exit...")
}
