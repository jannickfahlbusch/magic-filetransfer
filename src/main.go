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
	usage           = flag.Bool("usage", false, "Print the usage for magic-filetransfer and exit")
	help            = flag.Bool("help", false, "Print the usage for magic-filetransfer and exit - alias for '-usage'")
	version         = flag.Bool("version", false, "Print the version for magic-filetransfer")
)

func main() {
	flag.Parse()

	//Check, if the user wants to see the usage of mft
	if *usage || *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *version {
		fmt.Printf("----- magic-filetransfer (mft) -------\n")
		fmt.Printf("Jannick Fahlbusch <git@jf-projects.de>\n")
		fmt.Printf("           Version 0.1-alpha          \n")

		os.Exit(0)
	}

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
