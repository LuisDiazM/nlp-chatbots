package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type RequestCreateLicense struct {
	UserId      string `json:"user_id,omitempty"`
	LicenseType string `json:"license_type,omitempty"`
}

type RequestLicense struct {
	UserId string `json:"user_id,omitempty"`
}

func main() {
	// Connect al servidor NATS
	nc, err := nats.Connect("nats://nlpchatbots:random94566546@localhost:4222")
	if err != nil {
		log.Fatalf("Error al conectarse a NATS: %v", err)
	}
	defer nc.Close()

	// Definir el tema al que se enviará el mensaje
	subject := "license.create"
	subject2 := "license.get"

	requestData := RequestCreateLicense{UserId: "luismiguel@gmail.com", LicenseType: "FREE"}
	data, err := json.Marshal(requestData)
	if err != nil {
		log.Fatalln(err)
	}
	// Mensaje a enviar

	// Publicar el mensaje en el tema
	err = nc.Publish(subject, data)
	if err != nil {
		log.Fatalf("Error al publicar el mensaje: %v", err)
	}
	fmt.Println("Mensaje publicado exitosamente.")

	requestreply := RequestLicense{UserId: "luismiguel@gmail.com"}
	data2, err := json.Marshal(requestreply)
	if err != nil {
		log.Fatalln(err)
	}
	msf, err := nc.Request(subject2, data2, 10*time.Second)
	var sop interface{}
	json.Unmarshal(msf.Data, &sop)
	fmt.Println(sop)

}
