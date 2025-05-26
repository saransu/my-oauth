package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"
)

func InitSocket() {
	ln, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal(err.Error())
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}

		go handleSocketConn(conn)
	}
}

func handleSocketConn(conn net.Conn) {
	defer conn.Close()
	fmt.Println("connection established!")

	ticker := time.NewTicker(time.Second)
	for {
		_, err := conn.Write([]byte(fmt.Sprintf("sending from server: %d\n", rand.Int31())))

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		<-ticker.C
	}
}
