package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type transmitFileRequest struct {
	FileName string
	Size     int64
}

func broadcastFileTransfer(transmitRequest transmitFileRequest, c chan net.IP) {
	fmt.Printf("Offering %s to all members of the network...\n", transmitRequest.FileName)

	connection, err := net.DialUDP("udp4", nil, &net.UDPAddr{
		IP:   net.IPv4bcast,
		Port: 13159,
	})

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	//Build the request
	message, error := json.Marshal(transmitRequest)
	if err != nil {
		log.Fatal(error)
	}

	for {
		connection.Write([]byte(message))

	}
	//c <- net.IPv4(255, w255, 255, 255)
}

func startFileTransfer(peer net.IPAddr, fileName string) {

}
