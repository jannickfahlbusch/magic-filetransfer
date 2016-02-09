package main

import (
	"log"
	"os"
	"path/filepath"
)

type transmitFileRequest struct {
	FileName string
	Size     int64
}

func buildTransmitFileRequest(fileName string) (transmitFileRequest, error) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileNameStat, err := file.Stat()
	if err != nil {
		return transmitFileRequest{"", 0}, err
	}

	_, baseFileName := filepath.Split(fileName)
	transmitRequest := transmitFileRequest{baseFileName, fileNameStat.Size()}

	return transmitRequest, nil
}
