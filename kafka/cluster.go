package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
)

func main() {

	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	config.Consumer.Offsets.CommitInterval = 1 * time.Second
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	brokers := []string{"192.168.0.16:9092"}
	topics := []string{"test", "bidder-test"}
	consumer, err := cluster.NewConsumer(brokers, "my-group", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case err, more := <-consumer.Errors():
			if more {
				log.Printf("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				log.Printf("Rebalanced: %+v\n", ntf)
			}
		case <-signals:
			fmt.Println("aaa")
			return
		}
	}
}
