package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func getFileHash(filePath string) (string, error) {
	var result string

	file, error := os.Open(filePath)
	if error != nil {
		return result, error
	}

	defer file.Close()

	hash := md5.New()
	_, error = io.Copy(hash, file)
	if error != nil {
		return result, error
	}

	result = hex.EncodeToString(hash.Sum([]byte{})[:15])
	return result, nil
}

func compareHash(originalFileHash string, filePath string) bool {
	fileHash, error := getFileHash(filePath)
	if error != nil {
		log.Fatal(error)
	}

	return originalFileHash == fileHash
}
