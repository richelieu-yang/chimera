// Package rsaKit 支持长文本加解密
/*
Golang-RSA加密解密-数据无大小限制：https://www.cnblogs.com/akidongzi/p/12036165.html

密钥对（公钥、私钥）的要求：	PKCS#8、PEM
在线生成非对称加密公钥、私钥：	http://web.chacuo.net/netrsakeypair
*/
package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/crypto/base64Kit"
)

func EncryptWithPKCS1(data, publicKey []byte) ([]byte, error) {
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

func DecryptWithPKCS1(data, privateKey, password []byte) ([]byte, error) {
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
