package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

var ch chan int = make(chan int)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go reader(conn)

	reader := bufio.NewReader(os.Stdin)

	localAddr := conn.LocalAddr().String()

	for {
		data, _, _ := reader.ReadLine()
		b := []byte(localAddr + ": " + string(data))
		conn.Write(b)

		select {
		case <-ch:
			log.Fatal("read server error")
		default:

		}
	}
}

func reader(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			ch <- 1
			log.Fatal(err)
		}
		fmt.Println(string(buf))
	}
}
