package main

import (
	"encoding/json"
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
		length, remoteAddr, err := socket.ReadFromUDP(data)
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

		go acceptIncomingFileTransfer(socket, remoteAddr.IP, request)
		//c <- remoteAddr.IP
	}
}

func acceptIncomingFileTransfer(connection *net.UDPConn, peer net.IP, offering transmitFileRequest) {
	//@ToDo
	//Peer accept senden
	//Auf Datei warten
	//Datei empfangen
	//Datei speichern
}
