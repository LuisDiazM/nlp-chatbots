package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

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

	// Mensaje a enviar
	message := []byte("Hola, este es un mensaje de ejemplo!")

	// Publicar el mensaje en el tema
	err = nc.Publish(subject, message)
	if err != nil {
		log.Fatalf("Error al publicar el mensaje: %v", err)
	}
	err = nc.Publish(subject2, message)
	if err != nil {
		log.Fatalf("Error al publicar el mensaje: %v", err)
	}
	fmt.Println("Mensaje publicado exitosamente.")
}
