package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
	"github.com/richelieu42/chimera/v2/src/crypto/base64Kit"
)

// Encrypt 公钥加密
func Encrypt(data, publicKey []byte) ([]byte, error) {
	opts := loadOptions()
	return opts.Encrypt(data, publicKey)
}

func (opts *rsaOptions) Encrypt(data, publicKey []byte) ([]byte, error) {
	if err := sliceKit.AssertNotEmpty(publicKey); err != nil {
		return nil, err
	}

	// 公钥
	key, err := opts.parsePublicKey(publicKey)
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
	return base64Kit.Encode1(data), nil
}
