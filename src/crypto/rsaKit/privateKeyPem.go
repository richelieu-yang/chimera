package rsaKit

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

// EncryptPrivatePEM 通过password，加密私钥（.pem格式）
/*
encryptWithPKCS8 a pem private key
input: 	pem raw
output:	pem raw
*/
func (opts *rsaOptions) EncryptPrivatePEM(pemRaw []byte) ([]byte, error) {
	if strKit.IsEmpty(opts.password) {
		// 密码为空，不加密私钥
		return pemRaw, nil
	}

	block, _ := pem.Decode(pemRaw)
	if block == nil {
		return nil, errorKit.Simple("fail to decode pem because block is nil")
	}
	der := block.Bytes
	privateKey, err := derToPrivateKey(der)
	if err != nil {
		return nil, err
	}
	block, err = opts.encryptPrivateKey(privateKey)
	if err != nil {
		return nil, err
	}
	return pem.EncodeToMemory(block), nil
}

// encryptPrivateKey
/*
input: private key
output: pem block
*/
func (opts *rsaOptions) encryptPrivateKey(privateKey interface{}) (*pem.Block, error) {
	switch k := privateKey.(type) {
	case *ecdsa.PrivateKey:
		raw, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, err
		}
		block, err := x509.EncryptPEMBlock(rand.Reader, "EC PRIVATE KEY", raw, []byte(opts.password), x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}
		return block, nil
	case *rsa.PrivateKey:
		var raw []byte
		var err error

		switch opts.format {
		case PKCS1:
			raw = x509.MarshalPKCS1PrivateKey(k)
		case PKCS8:
			raw, err = x509.MarshalPKCS8PrivateKey(k)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errorKit.Simple("invalid format(%v)", opts.format)
		}
		block, err := x509.EncryptPEMBlock(rand.Reader, "PRIVATE KEY", raw, []byte(opts.password), x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}
		return block, nil
	default:
		return nil, errorKit.Simple("Invalid key type. It must be *ecdsa.PrivateKey or *rsa.PrivateKey")
	}
}

// DecryptPrivatePEM 通过password，解密私钥（.pem格式）
/*
input: 	pem raw
output: pem raw
*/
func (opts *rsaOptions) DecryptPrivatePEM(pemRaw []byte) ([]byte, error) {
	if strKit.IsEmpty(opts.password) {
		// 密码为空，不解密私钥
		return pemRaw, nil
	}

	block, _ := pem.Decode(pemRaw)
	if block == nil {
		return nil, errorKit.Simple("fail to decode pem because block is nil")
	}
	if !x509.IsEncryptedPEMBlock(block) {
		return nil, errorKit.Simple("fail to decode pem because it's not a decrypted pem")
	}
	der, err := x509.DecryptPEMBlock(block, []byte(opts.password))
	if err != nil {
		return nil, err
	}
	privateKey, err := derToPrivateKey(der)
	if err != nil {
		return nil, err
	}

	var raw []byte
	switch k := privateKey.(type) {
	case *ecdsa.PrivateKey:
		raw, err = x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, err
		}
	case *rsa.PrivateKey:
		switch opts.format {
		case PKCS1:
			raw = x509.MarshalPKCS1PrivateKey(k)
		case PKCS8:
			raw, err = x509.MarshalPKCS8PrivateKey(k)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errorKit.Simple("invalid format(%v)", opts.format)
		}
	default:
		return nil, errorKit.Simple("Invalid key type. It must be *ecdsa.PrivateKey or *rsa.PrivateKey")
	}

	rawBase64 := base64.StdEncoding.EncodeToString(raw)
	derBase64 := base64.StdEncoding.EncodeToString(der)
	if rawBase64 != derBase64 {
		return nil, errorKit.Simple("invalid PEM: raw does not match with der")
	}
	block = &pem.Block{
		Type:  block.Type,
		Bytes: der,
	}
	return pem.EncodeToMemory(block), nil
}

func derToPrivateKey(der []byte) (key interface{}, err error) {
	if key, err = x509.ParsePKCS1PrivateKey(der); err == nil {
		return key, nil
	}
	if key, err = x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key.(type) {
		case *rsa.PrivateKey, *ecdsa.PrivateKey:
			return
		default:
			return nil, errorKit.Simple("Found unknown private key type in PKCS#8 wrapping")
		}
	}
	if key, err = x509.ParseECPrivateKey(der); err == nil {
		return
	}
	return nil, errorKit.Simple("Invalid key type. The DER must contain an rsa.PrivateKey or ecdsa.PrivateKey")
}
