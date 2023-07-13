package rsaKit

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
)

// ParsePublicKeyFromPem 解析公钥.
/*
支持: PKCS1、PKCS8.
*/
func ParsePublicKeyFromPem(pemData []byte) (*rsa.PublicKey, error) {
	return jwt.ParseRSAPublicKeyFromPEM(pemData)

	//block, _ := pem.Decode(pemData)
	//if block == nil {
	//	return nil, errorKit.New("fail to decode pem because block is nil")
	//}
	//
	//keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	//if err != nil {
	//	return nil, err
	//}
	//return keyInterface.(*rsa.PublicKey), nil
}

// ParsePrivateKeyFromPem 解析私钥.
/*
支持: PKCS1、PKCS8.

@param password 私钥的密码.
*/
func ParsePrivateKeyFromPem(data []byte) (*rsa.PrivateKey, error) {
	return jwt.ParseRSAPrivateKeyFromPEM(data)

	//block, _ := pem.Decode(data)
	//if block == nil {
	//	return nil, errorKit.New("fail to decode pem because block is nil")
	//}
	//
	//switch opts.format {
	//case PKCS1:
	//	return x509.ParsePKCS1PrivateKey(block.Bytes)
	//case PKCS8:
	//	keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	//	if err != nil {
	//		return nil, err
	//	}
	//	return keyInterface.(*rsa.PrivateKey), nil
	//default:
	//	return nil, errorKit.New("invalid key format(%d)", opts.format)
	//}
}
