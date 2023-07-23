package app

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func CreateLicense(app *Application) {
	_, err := app.Nats.Conn.QueueSubscribe(subjectCreateLicense, queueLicenseManager, func(msg *nats.Msg) {
		fmt.Printf("Recibido mensaje en '%s': %s\n", msg.Subject, string(msg.Data))
	})
	if err != nil {
		log.Fatalf("Error al suscribirse: %v", err)
	}
}

func GetLicense(app *Application) {
	_, err := app.Nats.Conn.QueueSubscribe(subjectGetLicense, queueLicenseManager, func(msg *nats.Msg) {
		fmt.Printf("Recibido mensaje en '%s': %s\n", msg.Subject, string(msg.Data))
	})
	if err != nil {
		log.Printf("Error al suscribirse: %v", err)
	}
}

func UpdateLicense(app *Application) {
	_, err := app.Nats.Conn.QueueSubscribe(subjectUpdateLicense, queueLicenseManager, func(msg *nats.Msg) {
		fmt.Printf("Recibido mensaje en '%s': %s\n", msg.Subject, string(msg.Data))
	})
	if err != nil {
		log.Printf("Error al suscribirse: %v", err)
	}
}

func DeleteLicense(app *Application) {
	_, err := app.Nats.Conn.QueueSubscribe(subjectDeleteLicense, queueLicenseManager, func(msg *nats.Msg) {
		fmt.Printf("Recibido mensaje en '%s': %s\n", msg.Subject, string(msg.Data))
	})
	if err != nil {
		log.Printf("Error al suscribirse: %v", err)
	}
}
