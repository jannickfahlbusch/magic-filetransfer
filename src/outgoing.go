package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/cheggaaa/pb"
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

	go listenToStartFileTransfer(c)

	for {
		//Write the offer to broadcast
		connection.Write([]byte(message))

		//data := make([]byte, 4096)

		//c <- remoteAddr.IP
		time.Sleep(time.Second)
	}
	//c <- net.IPv4(255, w255, 255, 255)
}

func listenToStartFileTransfer(c chan net.IP) {
	connection, error := net.Listen("tcp4", ":13160")
	if error != nil {
		log.Fatal(error)
	}

	defer connection.Close()

	for {
		ln, error := connection.Accept()
		if error != nil {
			log.Fatal(error)
		}

		file, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		//Get a stat of the file and instanciate the progress-bar
		fileStat, _ := file.Stat()
		progressBar := pb.New64(fileStat.Size()).SetUnits(pb.U_BYTES)
		progressBar.Prefix(*fileName + ": ")
		progressBar.Start()

		fmt.Printf("Writing %s to %s\n", *fileName, ln.RemoteAddr().String())

		//Create a proxy for the file reader
		fileReaderProxy := progressBar.NewProxyReader(file)

		written, error := io.Copy(ln, fileReaderProxy)
		if error != nil {
			log.Fatal(error)
		}

		fmt.Printf("Written %v bytes to %s\n", written, ln.RemoteAddr().String())

		c <- net.ParseIP(ln.RemoteAddr().String())
	}
}
