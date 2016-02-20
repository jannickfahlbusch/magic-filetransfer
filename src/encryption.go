package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

func encrypt(key, text []byte) (ciphertext []byte, err error) {

	var block cipher.Block

	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}

	ciphertext = make([]byte, aes.BlockSize+len(string(text)))

	// iv =  initialization vector
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], text)

	return
}

func decrypt(key, ciphertext []byte) (plaintext []byte, err error) {

	var block cipher.Block

	if block, err = aes.NewCipher(key); err != nil {
		return
	}

	if len(ciphertext) < aes.BlockSize {
		err = errors.New("ciphertext too short")
		return
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(ciphertext, ciphertext)

	plaintext = ciphertext

	return
}

func encryptFile(filePath string, passphrase []byte) io.Reader {

	var ciphertext, plaintext []byte
	var err error

	plaintext, _ = ioutil.ReadFile(filePath)
	if ciphertext, err = encrypt(passphrase, plaintext); err != nil {
		log.Fatal(err)
	}

	return bytes.NewReader(ciphertext)
}

func decryptFile(filePath string, passphrase []byte) io.Reader {

	var ciphertext, plaintext []byte
	var err error

	fmt.Println("here")

	ciphertext, err = ioutil.ReadFile(filePath + ".crypt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("here")

	if plaintext, err = decrypt(passphrase, ciphertext); err != nil {
		log.Fatal(err)
	}
	fmt.Println("here")

	return bytes.NewReader(plaintext)
}
