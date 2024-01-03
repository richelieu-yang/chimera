package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// pkcs7Padding adds padding to the plaintext according to the PKCS7 standard
func pkcs7Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// pkcs7UnPadding removes padding from the ciphertext according to the PKCS7 standard
func pkcs7UnPadding(ciphertext []byte) []byte {
	length := len(ciphertext)
	unpadding := int(ciphertext[length-1])
	return ciphertext[:(length - unpadding)]
}

// Encrypt encrypts the plaintext with the given key using AES/CBC/PKCS7
func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = pkcs7Padding(plaintext, block.BlockSize())
	ciphertext := make([]byte, block.BlockSize()+len(plaintext))
	iv := ciphertext[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[block.BlockSize():], plaintext)
	return ciphertext, nil
}

// Decrypt decrypts the ciphertext with the given key using AES/CBC/PKCS7
func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < block.BlockSize() {
		return nil, fmt.Errorf("ciphertext too short")
	}
	iv := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]
	if len(ciphertext)%block.BlockSize() != 0 {
		return nil, fmt.Errorf("ciphertext is not a multiple of the block size")
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = pkcs7UnPadding(ciphertext)
	return ciphertext, nil
}

func main() {
	plaintext := []byte("%7B%22method%22%3A10021%7D")

	// 16 bytes key for AES-128
	key := []byte("00000yozo_config")

	fmt.Printf("Plaintext: %s\n", plaintext)
	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext (base64): %s\n", base64.StdEncoding.EncodeToString(ciphertext))

	//ciphertext = []byte("5b2d3131362c202d3130372c202d32332c202d3130392c202d39352c202d34332c2031312c2031312c203132372c202d3132382c2035342c2036312c202d372c202d32342c202d312c202d3132312c2032342c2032382c2032372c202d31322c2036362c2035302c202d38352c20352c202d39302c202d35352c202d31302c2037332c202d3130302c202d36322c202d33312c202d3130355d")

	decrypted, err := Decrypt(ciphertext, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted: %s\n", decrypted)
}
