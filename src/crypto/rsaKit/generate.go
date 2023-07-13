package rsaKit

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
)

// GenerateKeyFiles 生成: 公钥 && 私钥
/*
@param bits		512 ||1024 || 2048 || 3072 || 4096
@param priPath	私钥文件存放的位置
@param pubPath	公钥文件存放的位置
@param options 	可配置: format、password...
*/
func GenerateKeyFiles(bits int, priPath, pubPath string, options ...RsaOption) error {
	pri, pub, err := GenerateKeys(bits, options...)
	if err != nil {
		return err
	}

	if err := fileKit.WriteToFile(pri, priPath); err != nil {
		return err
	}
	if err := fileKit.WriteToFile(pub, pubPath); err != nil {
		return err
	}
	return nil
}

func GenerateKeys(bits int, options ...RsaOption) (pri []byte, pub []byte, err error) {
	opts := loadOptions(options...)
	return opts.GenerateKeyPair(bits)
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
