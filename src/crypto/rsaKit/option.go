package rsaKit

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
)

type (
	rsaOptions struct {
		// format 密钥格式（私钥）
		format KeyFormat
		// password 私钥密码（可以为""）
		password string
	}

	RsaOption func(opts *rsaOptions)
)

func loadOptions(options ...RsaOption) *rsaOptions {
	opts := &rsaOptions{
		format:   PKCS8,
		password: "",
	}
	for _, option := range options {
		option(opts)
	}

	// check
	switch opts.format {
	case PKCS1:
		fallthrough
	case PKCS8:
		// do nothing
	default:
		opts.format = PKCS8
	}

	return opts
}

// WithFormat 配置: 密钥格式（私钥）
func WithFormat(format KeyFormat) RsaOption {
	return func(opts *rsaOptions) {
		opts.format = format
	}
}

// WithPassword 配置: 私钥密码（可以为""）
func WithPassword(password string) RsaOption {
	return func(opts *rsaOptions) {
		opts.password = password
	}
}

func (opts *rsaOptions) GenerateKeyPair(bits int) (pri []byte, pub []byte, err error) {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	var derStream []byte
	switch opts.format {
	case PKCS1:
		derStream = x509.MarshalPKCS1PrivateKey(privateKey)
	case PKCS8:
		derStream, err = x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, errorKit.New("invalid format(%v)", opts.format)
	}
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	pri = pem.EncodeToMemory(block)
	pri, err = opts.EncryptPrivatePEM(pri)
	if err != nil {
		return nil, nil, err
	}

	// 生成公钥
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return nil, nil, err
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	pub = pem.EncodeToMemory(block)

	return pri, pub, nil
}
