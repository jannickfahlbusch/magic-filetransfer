package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/cheggaaa/pb"
	"github.com/dustin/go-humanize"
)

var filePath string

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

	//@ToDO: Check, if the user wants the file
	var request transmitFileRequest
	error := json.Unmarshal(data[:length], &request)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Got connection attempt from %s -> ", remoteAddr)

	//Only accept connections from server when the user specified it
	if *server != nil {
		if remoteAddr.IP.Equal(*server) {
			fmt.Printf("accepting connection\n")
			fmt.Printf("Remote wants to transfer %s (Size: %s)\n", request.FileName, humanize.Bytes(uint64(request.Size)))
			acceptIncomingFileTransfer(remoteAddr.IP, request, c)
		} else {
			fmt.Printf("DENY (only accepting connections from %s)", *server)
		}
	} else {
		fmt.Printf("accepting connection\n")
		fmt.Printf("Remote wants to transfer %s (Size: %s)\n", request.FileName, humanize.Bytes(uint64(request.Size)))
		acceptIncomingFileTransfer(remoteAddr.IP, request, c)
	}
}

func acceptIncomingFileTransfer(peerAddress net.IP, offering transmitFileRequest, c chan net.IP) {
	remoteAddressHostPort := peerAddress.String() + ":13160"
	connection, err := net.Dial("tcp4", remoteAddressHostPort)
	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()

	fmt.Println("Connection established")

	if *outputDirectory != "" {
		filePath = *outputDirectory + "/"
	}

	if *outputFileName != "" {
		filePath += *outputFileName
	} else {
		filePath += offering.FileName
	}

	fmt.Printf("Saving %s from %s to %s\n", offering.FileName, peerAddress.String(), filePath)

	file, error := os.Create(filePath)
	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()

	progressBar := pb.New64(offering.Size).SetUnits(pb.U_BYTES)
	progressBar.Prefix(offering.FileName + ": ")
	progressBar.Start()

	//Create the writer proxy
	writerProxy := io.MultiWriter(file, progressBar)

	written, error := io.CopyN(writerProxy, connection, offering.Size)
	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("Recieved '%s' (Size: %s) from %s \n", offering.FileName, humanize.Bytes(uint64(written)), connection.RemoteAddr().String())
	if !compareHash(offering.Hash, filePath) {
		log.Fatal("Hash mismatch: The file was not transferred correctly!")
	}

	c <- peerAddress
}
