package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare("szy", false, false, false, false, nil)
	failOnError(err, "failed to declare a queue")

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	failOnError(err, "failed to register a consumer")

	var forever chan bool = make(chan bool, 0)

	go func() {
		for body := range msgs {
			fmt.Printf("received a message:%s", body.Body)
		}
	}()
	log.Println("press Ctrl+c to exit!")
	<-forever
}
