package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
)

var (
	fileName        = flag.String("fileName", "", "Name of the file to transmit")
	outputDirectory = flag.String("outputDir", "", "Directory to save the retrieved file")
	outputFileName  = flag.String("outputName", "", "Save the recieved file under a different name")
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

		_, fileName := filepath.Split(*fileName)
		transmitRequest := transmitFileRequest{fileName, fileNameStat.Size()}

		go broadcastFileTransfer(transmitRequest, channel)
	}

	ipAddress := <-channel
	fmt.Println(ipAddress.String())
}
