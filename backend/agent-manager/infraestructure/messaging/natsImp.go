package messaging

import (
	"fmt"
	"log"

	"github.com/LuisDiazM/agent-manager/cmd/config"
	"github.com/nats-io/nats.go"
)

type NatsImp struct {
	Conn *nats.Conn
}

func NewNatsImplementation(env *config.Env) *NatsImp {
	url := fmt.Sprintf(`nats://%s:%s`, env.NATS_URL, env.NATS_PORT)
	log.Println(url)
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatalln(err)
	}
	return &NatsImp{Conn: nc}
}

func (natsImp *NatsImp) CloseConnection() {
	natsImp.Conn.Close()
}
