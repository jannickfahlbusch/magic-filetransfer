package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"
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

	go listenToStartFileTransfer(connection)

	for {
		//Write the offer to broadcast
		connection.Write([]byte(message))

		//data := make([]byte, 4096)

		//c <- remoteAddr.IP
		time.Sleep(time.Second)
	}
	//c <- net.IPv4(255, w255, 255, 255)
}

func listenToStartFileTransfer(connection *net.UDPConn) {
	for {
		data := make([]byte, 4096)
		length, remoteAddr, err := connection.ReadFromUDP(data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Remote peer %s wants to transfer %s\n", remoteAddr, data)

		//@ToDO: Check, if the user wants the file
		var request transmitFileRequest
		error := json.Unmarshal(data[:length], &request)
		if error != nil {
			log.Fatal(error)
		}

		go startFileTransfer(connection, remoteAddr.IP, request)
	}
}

func startFileTransfer(connection *net.UDPConn, peerIPAddress net.IP, request transmitFileRequest) {

}
