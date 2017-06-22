package main

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
)

var (
	commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	key_text = "yoyayoyayoayasdasdfadsfbk11ufqwe"
)

func Encrypter(plaintext []byte) []byte {
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		log.Fatalln(err)
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	log.Printf("%s->%x\n", plaintext, ciphertext)
	return ciphertext
}

func Decrypter(plaintext []byte) {
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		log.Fatalln(err)
	}

	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, plaintext)
	log.Printf("%x=>%s\n", plaintext, plaintextCopy)
}

func main() {
	ciphertest := Encrypter([]byte("asdasdadas"))
	Decrypter(ciphertest)
}
