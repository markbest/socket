package main

import (
	"log"
	"net"
	. "github.com/markbest/socket/conf"
	"github.com/markbest/socket/protocol"
	"time"
)

func main() {
	if err := InitConfig(""); err != nil {
		panic(err)
	}

	netListen, err := net.Listen("tcp", Conf.App.Port)
	if err != nil {
		panic(err)
	}
	defer netListen.Close()

	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	tmpBuffer := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	go heartBeating(conn, readerChannel)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			return
		}
		tmpBuffer = protocol.Unpack(append(tmpBuffer, buffer[:n]...), readerChannel)
	}
}

//心跳计时，如果接收到消息，则重置连接超时时间
func heartBeating(conn net.Conn, readerChannel chan []byte) {
	for {
		select {
		case data := <-readerChannel:
			log.Println(string(data))
			conn.SetDeadline(time.Now().Add(time.Duration(Conf.App.Timeout) * time.Second))
		case <-time.After(time.Second * time.Duration(Conf.App.Timeout)):
			log.Println("It's really weird to get Nothing!!!")
			conn.Close()
		}
	}
}