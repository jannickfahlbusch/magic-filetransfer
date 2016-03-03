package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
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
		fmt.Printf("           Version 0.2          \n")

		os.Exit(0)
	}

	channel := make(chan net.IP)
	if *fileName == "" {
		fmt.Println("No filename given. Assuming, you want to recieve a file")

		listenForIncomingFileTransfer(channel)
	} else {
		fmt.Printf("Starting to offer %s to all other devices in the network...\n", *fileName)

		transmitRequest, error := buildTransmitFileRequest(*fileName)
		if error != nil {
			log.Fatal(error)
		}

		broadcastFileTransfer(transmitRequest, channel)
	}

	//Let main() wait until the transfer is completed
	<-channel
}
