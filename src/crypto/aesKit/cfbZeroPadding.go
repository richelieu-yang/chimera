package aesKit

import "github.com/gogf/gf/v2/crypto/gaes"

// EncryptCfbZeroPadding
/*
加密模式: CFB
填充方式: ZeroPadding
*/
var EncryptCfbZeroPadding func(plainData []byte, key []byte, padding *int, iv ...[]byte) ([]byte, error) = gaes.EncryptCFB

// DecryptCfbZeroPadding
/*
加密模式: CFB
填充方式: ZeroPadding
*/
var DecryptCfbZeroPadding func(cipherData []byte, key []byte, unPadding int, iv ...[]byte) ([]byte, error) = gaes.DecryptCFB
