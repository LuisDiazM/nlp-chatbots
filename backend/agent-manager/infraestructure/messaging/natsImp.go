package messaging

import (
	"log"

	"github.com/LuisDiazM/agent-manager/cmd/config"
	"github.com/nats-io/nats.go"
)

type NatsImp struct {
	Conn *nats.Conn
}

func NewNatsImplementation(env *config.Env) *NatsImp {
	nc, err := nats.Connect(env.NATS_URL)
	if err != nil {
		log.Fatalln(err)
	}
	return &NatsImp{Conn: nc}
}

func (natsImp *NatsImp) CloseConnection() {
	natsImp.Conn.Close()
}
