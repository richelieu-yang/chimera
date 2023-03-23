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
	"github.com/richelieu42/chimera/src/crypto/base64Kit"
	"github.com/richelieu42/chimera/src/resources"
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

// Encrypt 公钥加密
func Encrypt(publicKey []byte, data []byte) ([]byte, error) {
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

func EncryptToString(publicKeyData []byte, data []byte) (string, error) {
	data, err := Encrypt(publicKeyData, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Decrypt 私钥解密
/**
 * @param password 私钥的密码，可以为nil
 */
func Decrypt(privateKey, password, data []byte) ([]byte, error) {
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

// DecryptToString
/**
 * @param password 私钥的密码，可以为nil
 */
func DecryptToString(privateKeyData, password, data []byte) (string, error) {
	data, err := Decrypt(privateKeyData, password, data)
	if err != nil {
		return "", err
	}
	return string(data), nil
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
