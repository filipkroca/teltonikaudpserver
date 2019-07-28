package main

import (
	"fmt"
	"net"
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
func onUDPMessage(udpc *net.UDPConn, buf *[]byte, len int, addr *net.UDPAddr) {
	//conn := *udpc

	fmt.Println((*buf)[:len])
	(*udpc).WriteToUDP([]byte("hello world"), addr)

}
