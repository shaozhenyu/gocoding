package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/Shopify/sarama"
)

func main() {
	producer, err := sarama.NewAsyncProducer([]string{"192.168.0.16:9092"}, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
		log.Println("kkk")
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var errors int

ProducerLoop:
	for {
		select {
		case producer.Input() <- &sarama.ProducerMessage{Topic: "test", Key: nil, Value: sarama.StringEncoder("testing 123")}:
			log.Println("sss")
		case err := <-producer.Errors():
			log.Println("Failed to produce message", err)
			errors++
		case <-signals:
			log.Println("bbb")
			break ProducerLoop
		}
	}

	//log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
	log.Fatal("ttt")
}
