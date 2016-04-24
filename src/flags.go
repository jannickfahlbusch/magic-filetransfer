package main

import "gopkg.in/alecthomas/kingpin.v2"

var (
	fileName              = kingpin.Flag("fileName", "Name of the file to transmit").Short('f').ExistingFile()
	outputDirectory       = kingpin.Flag("outputDir", "Directory to save the retrieved file in").Short('o').ExistingDir()
	outputFileName        = kingpin.Flag("outputName", "Save the recieved file under a different name").String()
	usage                 = kingpin.Flag("usage", "Print the usage for magic-filetransfer and exit").Short('u').Short('h').Bool()
	version               = kingpin.Flag("version", "Print the version for magic-filetransfer").Short('v').Bool()
	client                = kingpin.Flag("client", "IP-Address of the client you want to send the file to").Short('c').IP()
	server                = kingpin.Flag("server", "IP-Address of the server you want to recieve the file from").Short('s').IP()
	disableIntegrityCheck = kingpin.Flag("disable-integrity-check", "Disable the integrity-checks after the transfer - NOT recommended!").Bool()
	daemon                = kingpin.Flag("daemon", "Run as daemon").Bool()
)
