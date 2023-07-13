package rsaKit

import (
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
