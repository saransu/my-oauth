package server

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Conn *amqp.Connection
var MainQueue amqp.Queue
var MainChannel *amqp.Channel

func InitQueue() {
	Conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err.Error())
	}

	MainChannel, err = Conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	// defer MainChannel.Close()

	MainQueue, err = MainChannel.QueueDeclare(
		"main-queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
