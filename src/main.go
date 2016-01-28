package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

var (
	fileName = flag.String("fileName", "", "Name of the file to transmit")
)

func main() {
	flag.Parse()

	channel := make(chan net.IP)
	if *fileName == "" {
		fmt.Println("No filename given. Assuming, you want to recieve a file")

		go listenForIncomingFileTransfer(channel)
	} else {

		file, err := os.Open(*fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		fileNameStat, err := file.Stat()
		if err != nil {
			return
		}

		transmitRequest := transmitFileRequest{*fileName, fileNameStat.Size()}

		go broadcastFileTransfer(transmitRequest, channel)
	}
	fmt.Println(<-channel)
}
