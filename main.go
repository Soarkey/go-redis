package main

import (
	"log"
	"net"

	"go-redis/tcp"
)

func main() {
	closeChan := make(chan struct{})
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println(err)
	}
	log.Printf("server %v start...\n", listener.Addr().String())
	tcp.ListenAndServe(listener, tcp.MakeEchoHandler(), closeChan)
}
