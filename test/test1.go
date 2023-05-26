package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
	"os"
)

// GenerateKeys 生成公钥私钥
/*
golang 生成RSA公钥和私钥文件 https://blog.csdn.net/lff1123/article/details/127788928

@param bits 		512 || 1024 || 2048 || 3072 || 4096
@param keyFormat	rsaKit.PKCS1 || rsaKit.PKCS8
*/
func GenerateKeys(bits int, keyFormat rsaKit.KeyFormat, privateKeyPassword string) error {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	var derStream []byte
	switch keyFormat {
	case rsaKit.PKCS1:
		derStream = x509.MarshalPKCS1PrivateKey(privateKey)
	case rsaKit.PKCS8:
		derStream, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return err
		}
	default:
		return errorKit.Simple("invalid keyFormat(%v)", keyFormat)
	}
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	file, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	file, err = os.Create("public.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(file, block)
	if err != nil {
		return err
	}
	return nil
}
