package desKit

import "github.com/gogf/gf/v2/crypto/gdes"

var EncryptECB func(plainText []byte, key []byte, padding int) ([]byte, error) = gdes.EncryptECB

var DecryptECB func(cipherText []byte, key []byte, padding int) ([]byte, error) = gdes.DecryptECB
