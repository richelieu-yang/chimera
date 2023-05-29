package rsaKit

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
)

// parsePublicKey 解析公钥
func (opts *rsaOptions) parsePublicKey(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errorKit.Simple("fail to decode pem because block is nil")
	}

	keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return keyInterface.(*rsa.PublicKey), nil
}

// parsePrivateKey 解析私钥
func (opts *rsaOptions) parsePrivateKey(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errorKit.Simple("fail to decode pem because block is nil")
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
		return nil, errorKit.Simple("invalid key format(%d)", opts.format)
	}
}
