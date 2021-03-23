package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func reader(sck io.Reader) {
	buf := make([]byte, 1024)
	for {
		msgLen, err := sck.Read(buf[:])
		if err != nil {
			return
		}
		println("CLIENT: received back:", string(buf[0:msgLen]))
	}
}

func main() {
	sck, err := net.Dial("unix", "../server/echo.sock")
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
	defer func() {
		sck.Close()
		fmt.Println("CLIENT: socket closed")
	}()

	go reader(sck)
	for {
		_, err := sck.Write([]byte("hi"))
		if err != nil {
			log.Fatal("CLIENT: write error:", err)
			break
		}
		time.Sleep(1e9)
	}
}
