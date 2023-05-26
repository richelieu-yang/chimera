// Package rsaKit 支持长文本加解密
/*
Golang-RSA加密解密-数据无大小限制：https://www.cnblogs.com/akidongzi/p/12036165.html
*/
package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/crypto/base64Kit"
	"github.com/richelieu42/chimera/v2/src/resources"
	"github.com/sirupsen/logrus"
)

var (
	// DefaultPublicKey 默认公钥
	DefaultPublicKey []byte

	// DefaultPrivateKey 默认私钥
	DefaultPrivateKey []byte

	// DefaultPassword 默认私钥的密码
	DefaultPassword = []byte("Y3l5")
)

func init() {
	var err error
	var path string

	path = "resources/crypto/rsa/pub.pem"
	DefaultPublicKey, err = resources.Asset(path)
	if err != nil {
		logrus.Fatal(err)
	}

	path = "resources/crypto/rsa/pri.pem"
	DefaultPrivateKey, err = resources.Asset(path)
	if err != nil {
		logrus.Fatal(err)
	}
}

// EncryptWithPKCS8 （公钥）加密
/*
PS: 支持大文本（内部分块加解密）.

@param data 		原文
@param publicKey 	公钥
*/
func EncryptWithPKCS8(data, publicKey []byte) ([]byte, error) {
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
	chunks := split(data, partLen)
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

// DecryptWithPKCS8 （私钥）解密
/*
PS: 支持大文本（内部分块加解密）.

@param data	 		密文
@param privateKey 	私钥
@param password 	私钥的密码（没有则传nil）
*/
func DecryptWithPKCS8(data, privateKey, password []byte) ([]byte, error) {
	if err := sliceKit.AssertNotEmpty(privateKey); err != nil {
		return nil, err
	}

	// 私钥
	if password != nil {
		var err error
		privateKey, err = DecryptPEM(privateKey, password)
		if err != nil {
			return nil, err
		}
	}
	key, err := parsePKCS8PrivateKey(privateKey)
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
	chunks := split(data, partLen)
	for _, chunk := range chunks {
		tmp, err := rsa.DecryptPKCS1v15(rand.Reader, key, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	return buffer.Bytes(), nil
}
