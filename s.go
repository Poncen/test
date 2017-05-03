package main

import (
	"fmt"
	"net"
)

var port string

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:10000")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("start listen")
	for {	
		fmt.Println("sd")
		tcpConn, _ := tcpListener.AcceptTCP()
		defer tcpConn.Close()
		fmt.Println("连接的客服端信息:", tcpConn.RemoteAddr().String())
	}
}
