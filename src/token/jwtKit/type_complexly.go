package jwtKit

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu-yang/chimera/v3/src/crypto/caesarKit"
)

func (j *JWT) NewComplexly(method jwt.SigningMethod, claims jwt.MapClaims, options ...jwt.TokenOption) (cipherText string, err error) {
	var tokenString string
	tokenString, err = j.New(method, claims, options...)
	if err != nil {
		return
	}

	cipherText = caesarKit.Encrypt(tokenString, j.shift)
	return
}

func (j *JWT) ParseComplexly(cipherText string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (jwt.MapClaims, error) {
	tokenString := caesarKit.Decrypt(cipherText, j.shift)

	return j.Parse(tokenString, keyFunc, options...)
}
