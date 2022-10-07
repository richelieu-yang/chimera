package rsaKit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
)

/**
* PKCS1
 */
func parsePKCS1PrivateKey(s []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(s)
	if block == nil {
		return nil, errorKit.Simple("private key error(%s)", "PKCS1")
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

/**
 * PKCS8
 */
func parsePKCS8PrivateKey(s []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(s)
	if block == nil {
		return nil, errorKit.Simple("private key error(%s)", "PKCS8")
	}

	keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return keyInterface.(*rsa.PrivateKey), nil
}
