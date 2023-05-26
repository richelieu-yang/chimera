package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/fileKit"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
)

func main() {
	pri, pub, err := GenerateKeys(1024, rsaKit.PKCS8, "cyy")
	if err != nil {
		panic(err)
	}
	if err := fileKit.WriteToFile(pri, "pri.pem"); err != nil {
		panic(err)
	}
	if err := fileKit.WriteToFile(pub, "pub.pem"); err != nil {
		panic(err)
	}
}

// GenerateKeys 生成公钥私钥
/*
golang 生成RSA公钥和私钥文件 https://blog.csdn.net/lff1123/article/details/127788928

@param bits 		512 || 1024 || 2048 || 3072 || 4096
@param keyFormat	rsaKit.PKCS1 || rsaKit.PKCS8
*/
func GenerateKeys(bits int, keyFormat rsaKit.KeyFormat, privateKeyPassword string) (pri []byte, pub []byte, err error) {
	// 生成私钥文件
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	var derStream []byte
	switch keyFormat {
	case rsaKit.PKCS1:
		derStream = x509.MarshalPKCS1PrivateKey(privateKey)
	case rsaKit.PKCS8:
		derStream, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errorKit.Simple("invalid keyFormat(%v)", keyFormat)
	}
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	pri = pem.EncodeToMemory(block)
	pri, err = rsaKit.EncryptPEM(pri, []byte(privateKeyPassword))
	if err != nil {
		return nil, nil, err
	}

	// 生成公钥文件
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pub = pem.EncodeToMemory(block)

	return pri, pub, nil
}
