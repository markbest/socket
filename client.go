package main

import (
	"fmt"
	"net"
	. "github.com/markbest/socket/conf"
	"github.com/markbest/socket/protocol"
)

func sender(conn net.Conn) {
	fmt.Println("connect success")
	for i := 0; i < 100; i++ {
		words := "{\"Id\":1,\"Name\":\"golang\",\"Message\":\"message\"}"
		conn.Write(protocol.Packet([]byte(words)))
	}
	fmt.Println("send over")
}

func main() {
	if err := InitConfig(""); err != nil {
		panic(err)
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp4", Conf.App.Port)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	sender(conn)
}
