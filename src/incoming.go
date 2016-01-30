package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func listenForIncomingFileTransfer(c chan net.IP) {
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 13159,
	})

	if err != nil {
		log.Fatal(err)
	}

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

	channel := make(chan net.IP)
	acceptIncomingFileTransfer(remoteAddr.IP, request, channel)
}

func acceptIncomingFileTransfer(peerAddress net.IP, offering transmitFileRequest, c chan net.IP) {
	remoteAddressHostPort := peerAddress.String() + ":13160"
	connection, err := net.Dial("tcp4", remoteAddressHostPort)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection established")

	file, error := os.Create(offering.FileName)
	if error != nil {
		log.Fatal(error)
	}

	written, error := io.Copy(file, connection)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Recieved '%s' (Size: %d) from %s \n", offering.FileName, written, connection.RemoteAddr().String())

	// Close all of the handlers!
	defer file.Close()
	defer connection.Close()

	c <- net.ParseIP(connection.RemoteAddr().String())
}
