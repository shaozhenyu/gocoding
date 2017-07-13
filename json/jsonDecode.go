package main

import (
	"fmt"
	"strings"
	"io"
	"log"
	"encoding/json"
)

func main() {
	const jsonStream = `{"Name":"szy", "Text":"I love you"}
						{"Name":"lyy", "Text":"love you too"}
	`

	type Message struct {
		Name, Text string
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		err := dec.Decode(&m)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
