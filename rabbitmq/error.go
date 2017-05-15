package rabbitmq

import (
	"fmt"
	"log"
)

func FailOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
