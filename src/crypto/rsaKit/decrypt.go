package rsaKit

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/crypto/base64Kit"
)

// Decrypt 私钥解密
/*
@param options 可配置: format、password
*/
func Decrypt(data, privateKey []byte, options ...RsaOption) ([]byte, error) {
	opts := loadOptions(options...)
	return opts.Decrypt(data, privateKey)
}

func (opts *rsaOptions) Decrypt(data, pemData []byte) ([]byte, error) {
	if err := sliceKit.AssertNotEmpty(pemData, "pemData"); err != nil {
		return nil, err
	}

	// 私钥
	pemData, err := opts.decryptPrivatePEM(pemData)
	if err != nil {
		return nil, err
	}
	privateKey, err := opts.ParsePrivateKeyFromPem(pemData)
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
	partLen := privateKey.N.BitLen() / 8
	chunks := sliceKit.Split(data, partLen)
	for _, chunk := range chunks {
		tmp, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(tmp)
	}
	return buffer.Bytes(), nil
}
