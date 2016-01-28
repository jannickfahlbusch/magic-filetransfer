package main

import (
	"fmt"
	"log"
	"net"
)

func listenForIncomingFileTransfer(c chan net.IP) {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 13159,
	})

	if err != nil {
		log.Fatal(err)
	}

	for {
		data := make([]byte, 4096)
		_, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Remote peer %s wants to transfer %s %T\n", remoteAddr, data, remoteAddr)
		c <- remoteAddr.IP
	}
}

func acceptIncomingFileTransfer(peer net.IPAddr, fileName string) {

}
