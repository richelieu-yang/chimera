// Package rsaKit 支持长文本加解密
/**
 * Golang-RSA加密解密-数据无大小限制：https://www.cnblogs.com/akidongzi/p/12036165.html
 *
 * 密钥对（公钥、私钥）的要求：	PKCS#8、PEM
 * 在线生成非对称加密公钥、私钥：	http://web.chacuo.net/netrsakeypair
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

// Encrypt （公钥）加密
/*
PS: 支持大文本（内部分块加解密）.

@param data 		原文
@param publicKey 	公钥
*/
func Encrypt(data, publicKey []byte) ([]byte, error) {
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
		tmp, err := encrypt(key, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	data = buffer.Bytes()

	// base64编码
	return base64Kit.Encode(data), nil
}

// Decrypt （私钥）解密
/*
PS: 支持大文本（内部分块加解密）.

@param data	 		密文
@param privateKey 	私钥
@param password 	私钥的密码（没有则传nil）
*/
func Decrypt(data, privateKey, password []byte) ([]byte, error) {
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
		tmp, err := decrypt(key, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	return buffer.Bytes(), nil
}

func encrypt(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
}

func decrypt(privateKey *rsa.PrivateKey, data []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, privateKey, data)
}

func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
