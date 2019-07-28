package main

import (
	"fmt"
	"log"
	"net"
)

//Server is exported and hold the port number
type Server struct {
	Protocol string
	IP       []byte
	Port     int
}

//New start listening on specified port, should provide callback function. On new connection invoke a callback function as a new goroutine
func (t *Server) New(callBack func(udpc *net.UDPConn, buf *[]byte, len int, addr *net.UDPAddr)) {

	udpc, err := net.ListenUDP(t.Protocol, &net.UDPAddr{IP: t.IP, Port: t.Port, Zone: ""})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer udpc.Close()

	fmt.Printf("Listening on %v\n", udpc.LocalAddr())

	for {
		//make a buffer
		buf := make([]byte, 4096)
		n, addr, err := udpc.ReadFromUDP(buf)
		if err != nil {
			log.Panic("error when listening ", err)
		}

		//slice data
		sliced := buf[:n]

		fmt.Printf("New connection from %v , fired a new goroutine \n", addr)
		//on connection fire new goroutine
		go callBack(udpc, &sliced, n, addr)
	}
}
