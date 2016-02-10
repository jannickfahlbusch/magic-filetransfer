package main

import (
	"log"
	"os"
	"path/filepath"
)

type transmitFileRequest struct {
	FileName string
	Size     int64
	Hash     string
}

func buildTransmitFileRequest(fileName string) (transmitFileRequest, error) {
	var fileHash string
	transmitRequest := transmitFileRequest{"", 0, fileName}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileNameStat, err := file.Stat()
	if err != nil {
		return transmitRequest, err
	}

	fileHash, error := getFileHash(fileName)
	if error != nil {
		return transmitRequest, error
	}

	_, baseFileName := filepath.Split(fileName)
	transmitRequest = transmitFileRequest{baseFileName, fileNameStat.Size(), fileHash}

	return transmitRequest, nil
}
