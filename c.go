package main

import (
	"fmt"
	"net"
)

var port string

func main() {
	fmt.Println("Input your server's listen port")
	fmt.Scanln(&port)
	tcpAddr, _ := net.ResolveTCPAddr("tcp", port)
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
	data := make([]byte, 10000)
	for {
		conn.Read(data)
	}
}
