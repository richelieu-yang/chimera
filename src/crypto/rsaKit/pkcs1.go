package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/crypto/base64Kit"
)

// PKCS1Encrypt 公钥加密（密钥格式: PKCS1）
func PKCS1Encrypt(data, publicKey []byte) ([]byte, error) {
	if err := sliceKit.AssertNotEmpty(publicKey); err != nil {
		return nil, err
	}

	// 公钥
	key, err := parsePublicKey(publicKey)
	if err != nil {
		return nil, err
	}

	// 分块加密
	buffer := bytes.NewBufferString("")
	partLen := key.N.BitLen()/8 - 11
	chunks := sliceKit.Split(data, partLen)
	for _, chunk := range chunks {
		tmp, err := rsa.EncryptPKCS1v15(rand.Reader, key, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	data = buffer.Bytes()

	// base64编码
	return base64Kit.Encode(data), nil
}

// PKCS1Decrypt 公钥解密（密钥格式: PKCS1）
func PKCS1Decrypt(data, privateKey, privateKeyPassword []byte) ([]byte, error) {
	if err := sliceKit.AssertNotEmpty(privateKey); err != nil {
		return nil, err
	}

	// 私钥
	if privateKeyPassword != nil {
		var err error
		privateKey, err = DecryptPEM(privateKey, privateKeyPassword)
		if err != nil {
			return nil, err
		}
	}
	key, err := parsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	// base64解码
	data, err = base64Kit.Decode(data)
	if err != nil {
		return nil, err
	}

	// 分块解密
	buffer := bytes.NewBufferString("")
	partLen := key.N.BitLen() / 8
	chunks := sliceKit.Split(data, partLen)
	for _, chunk := range chunks {
		tmp, err := rsa.DecryptPKCS1v15(rand.Reader, key, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	return buffer.Bytes(), nil
}
