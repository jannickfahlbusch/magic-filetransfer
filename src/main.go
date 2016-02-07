package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	fileName        = kingpin.Flag("fileName", "Name of the file to transmit").String()
	outputDirectory = kingpin.Flag("outputDir", "Directory to save the retrieved file in").String()
	outputFileName  = kingpin.Flag("outputName", "Save the recieved file under a different name").ExistingFile()
	usage           = kingpin.Flag("usage", "Print the usage for magic-filetransfer and exit").Short('u').Short('h').Bool()
	version         = kingpin.Flag("version", "Print the version for magic-filetransfer").Short('v').Bool()
)

func main() {
	kingpin.Parse()

	//Check, if the user wants to see the usage of mft
	if *usage {
		kingpin.Usage()
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
