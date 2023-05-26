package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/crypto/base64Kit"
)

// PKCS8Encrypt 公钥加密（密钥格式: PKCS8）
/*
@param data 		原文
@param publicKey 	公钥
*/
func PKCS8Encrypt(data, publicKey []byte) ([]byte, error) {
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

// PKCS8Decrypt 私钥解密（密钥格式: PKCS8）
/*
PS: 支持大文本（内部分块加解密）.

@param data	 		密文
@param privateKey 	私钥
@param password 	私钥的密码（没有则传nil）
*/
func PKCS8Decrypt(data, privateKey, password []byte) ([]byte, error) {
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
