package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type User struct {
	Name string
	Ip   string
	Conn net.Conn
}

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	user := map[string]*User{}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}

		ip := conn.RemoteAddr().String()
		fmt.Println("connect : ", conn.RemoteAddr().String())
		if _, ok := user[ip]; !ok {
			newOne := &User{
				Name: ip,
				Ip:   ip,
				Conn: conn,
			}
			user[ip] = newOne
		}
		go handleConnection(conn, user[ip], user)
	}
}

func handleConnection(conn net.Conn, user *User, all map[string]*User) {
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
		//send to someone
		for k, v := range all {
			if k != user.Ip {
				v.Conn.Write(buf)
			}
		}
		//fmt.Println(string(buf))
	}
}
