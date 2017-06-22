package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

func main() {

	brokers := []string{"192.168.0.16:9092", "192.168.0.17:9092", "192.168.0.18:9092"}
	topic := "test"

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitions, err := consumer.Partitions(topic)

	var wg sync.WaitGroup
	wg.Add(len(partitions))

	for _, partition := range partitions {
		log.Println("consume partition:", partition)

		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt)

		go func(partition int32) {

			partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
			if err != nil {
				log.Fatalln(err)
				return
			}
			defer func() {
				wg.Done()
				if err := partitionConsumer.Close(); err != nil {
					log.Fatalln(err)
				}
			}()

		ConsumerLoop:
			for {
				select {
				case err := <-partitionConsumer.Errors():
					if err != nil {
						log.Fatalln(err)
					}
				case msg := <-partitionConsumer.Messages():
					if msg != nil {
						log.Println("topic:", msg.Topic, "partition:", msg.Partition, "offset:", msg.Offset, "message:", string(msg.Value))
						//do something
					}
				case <-shutdown:
					log.Println("stop consuming partition:", partition)
					break ConsumerLoop
				}
			}
		}(partition)
	}

	wg.Wait()

}
