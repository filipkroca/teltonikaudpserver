package main

import (
	"fmt"
	"log"
	"net"

	"github.com/filipkroca/teltonikaparser"
)

func main() {

	server := Server{
		Protocol: "udp",
		IP:       []byte{0, 0, 0, 0},
		Port:     49152,
	}

	//create new server
	server.New(onUDPMessage)
	defer fmt.Println("server closed")

}

//onUDPMessage is invoked when packet arrive
func onUDPMessage(udpc *net.UDPConn, dataBs *[]byte, len int, addr *net.UDPAddr) {
	//conn := *udpc

	x, err := teltonikaparser.Decode(dataBs)
	if err != nil {
		log.Panic("Unable to decode packet", err)
	}

	fmt.Printf("%+v", x)

	(*udpc).WriteToUDP([]byte("hello world"), addr)

}
