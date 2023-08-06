package app

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/LuisDiazM/nlp-chatbots/licensing-manager/domain/licenses/entities"
	"github.com/nats-io/nats.go"
)

func CreateLicense(app *Application) {
	_, err := app.Nats.Conn.QueueSubscribe(subjectCreateLicense, queueLicenseManager, func(msg *nats.Msg) {
		fmt.Printf("message '%s': %s\n", msg.Subject, string(msg.Data))
		var requestData RequestCreateLicense
		err := json.Unmarshal(msg.Data, &requestData)
		if err != nil {
			log.Println(err)
		}
		timestampNow := time.Now()
		monthsFreeLicenseAlive := 1
		licenseExpire := timestampNow.AddDate(0, monthsFreeLicenseAlive, 0)
		var license entities.License = entities.License{
			ExpiredAt: licenseExpire,
			UserId:    requestData.UserId,
			CreatedAt: timestampNow,
			Type:      requestData.LicenseType,
			Features:  entities.LicenseFeature{RateLimit: 1000, Trainings: 10},
		}
		result := app.LicenseUsecase.CreateLicense(license, context.TODO())
		log.Println(*result)
	})
	if err != nil {
		log.Fatalf("Error al suscribirse: %v", err)
	}
}

func GetLicense(app *Application) {
	_, err := app.Nats.Conn.QueueSubscribe(subjectGetLicense, queueLicenseManager, func(msg *nats.Msg) {
		fmt.Printf("message '%s': %s\n", msg.Subject, string(msg.Data))
		var requestData RequestLicense
		err := json.Unmarshal(msg.Data, &requestData)
		if err != nil {
			log.Println(err)
		}
		licenses := app.LicenseUsecase.GetLicensesByUserId(requestData.UserId, context.TODO())
		data, err := json.Marshal(licenses)
		if err != nil {
			log.Println(err)
		}
		err = msg.Respond(data)
		if err != nil {
			log.Println(err)
		}
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
