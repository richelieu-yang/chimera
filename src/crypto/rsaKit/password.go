// Package rsaKit 对私钥进行"加密"、"解密"操作.
/**
 * 参考（将里面两个PKCS1 marshall改为PKCS8的了）：https://www.jianshu.com/p/c102a639cc50
 */
package rsaKit

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/richelieu42/go-scales/src/core/errorKit"
)

// EncryptPEM 通过password，加密私钥
/**
 * encrypt a pem private key
 * input: pem raw
 * output: pem raw
 */
func EncryptPEM(pemRaw []byte, passwd []byte) ([]byte, error) {
	block, _ := pem.Decode(pemRaw)
	if block == nil {
		return nil, errorKit.Simple("fail to decode pem because block is nil")
	}

	der := block.Bytes

	privateKey, err := derToPrivateKey(der)
	if err != nil {
		return nil, err
	}

	block, err = encryptPrivateKey(privateKey, passwd)
	if err != nil {
		return nil, err
	}

	return pem.EncodeToMemory(block), nil
}

// encrypt a private key
// input: private key
// output: pem block
func encryptPrivateKey(privateKey interface{}, passwd []byte) (*pem.Block, error) {
	switch k := privateKey.(type) {
	case *ecdsa.PrivateKey:
		raw, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			return nil, err
		}

		block, err := x509.EncryptPEMBlock(rand.Reader, "EC PRIVATE KEY", raw, passwd, x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}

		return block, nil
	case *rsa.PrivateKey:
		raw, err := x509.MarshalPKCS8PrivateKey(k)
		if err != nil {
			return nil, err
		}

		block, err := x509.EncryptPEMBlock(rand.Reader, "RSA PRIVATE KEY", raw, passwd, x509.PEMCipherAES256)
		if err != nil {
			return nil, err
		}

		return block, nil
	default:
		return nil, errorKit.Simple("Invalid key type. It must be *ecdsa.PrivateKey or *rsa.PrivateKey")
	}
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

// DecryptPEM 通过password，解密私钥
/**
 * decrypt a pem private key
 * input: pem raw
 * output: pem raw
 */
func DecryptPEM(pemRaw []byte, passwd []byte) ([]byte, error) {
	block, _ := pem.Decode(pemRaw)
	if block == nil {
		return nil, errorKit.Simple("fail to decode pem because block is nil")
	}

	if !x509.IsEncryptedPEMBlock(block) {
		return nil, errorKit.Simple("fail to decode pem because it's not a decrypted pem")
	}

	der, err := x509.DecryptPEMBlock(block, passwd)
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
		raw, err = x509.MarshalPKCS8PrivateKey(k)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errorKit.Simple("Invalid key type. It must be *ecdsa.PrivateKey or *rsa.PrivateKey")
	}

	rawBase64 := base64.StdEncoding.EncodeToString(raw)
	derBase64 := base64.StdEncoding.EncodeToString(der)
	if rawBase64 != derBase64 {
		return nil, errorKit.Simple("Invalid decrypt PEM: raw does not match with der")
	}

	block = &pem.Block{
		Type:  block.Type,
		Bytes: der,
	}

	return pem.EncodeToMemory(block), nil
}
