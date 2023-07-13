package rsaKit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

// ParsePublicKeyFromPem 解析公钥
func (opts *rsaOptions) ParsePublicKeyFromPem(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errorKit.New("fail to decode pem because block is nil")
	}

	keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return keyInterface.(*rsa.PublicKey), nil
}

// ParsePrivateKeyFromPem 解析私钥
func (opts *rsaOptions) ParsePrivateKeyFromPem(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errorKit.New("fail to decode pem because block is nil")
	}

	switch opts.format {
	case PKCS1:
		return x509.ParsePKCS1PrivateKey(block.Bytes)
	case PKCS8:
		keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		return keyInterface.(*rsa.PrivateKey), nil
	default:
		return nil, errorKit.New("invalid key format(%d)", opts.format)
	}
}
