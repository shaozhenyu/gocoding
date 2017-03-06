package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("connect : ", conn.RemoteAddr().String())

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var ch chan int = make(chan int)
	r := bufio.NewReader(os.Stdin)
	go func() {
		for {
			data, _, _ := r.ReadLine()
			b := []byte("server say : " + string(data))
			conn.Write(b)

			select {
			case <-ch:
				log.Fatal("read client error")
			default:
			}
		}
	}()
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
