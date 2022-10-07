package rsaKit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu42/go-scales/src/core/errorKit"
)

func parsePublicKey(s []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(s)
	if block == nil {
		return nil, errorKit.Simple("public key error")
	}

	keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return keyInterface.(*rsa.PublicKey), nil
}
