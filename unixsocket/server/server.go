package main

import (
	"fmt"
	"net"
)

func echoServer(sckConn net.Conn) {
	for {
		buf := make([]byte, 512)
		msgLen, err := sckConn.Read(buf)
		if err != nil {
			println("SERVER: error1")
			return
		}

		data := buf[0:msgLen]
		data = append(data, '!')
		fmt.Printf("SERVER: Received: %v", string(data))
		_, err = sckConn.Write(data)
		if err != nil {
			panic("SERVER: Write error") // err.String()
		}
	}
}

func main() {
	sck, err := net.Listen("unix", "./echo.sock")
	if err != nil {
		println("SERVER: listen error") //, err.String())
		return
	}
	defer func() {
		sck.Close()
		fmt.Println("SERVER: socket closed")
	}()

	println("SERVER: listening")
	for {
		sckConn, err := sck.Accept()
		if err != nil {
			println("SERVER: accept error") //, err.String())
			return
		}

		go echoServer(sckConn)
	}
}
