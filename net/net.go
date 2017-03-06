package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(conn)
}
