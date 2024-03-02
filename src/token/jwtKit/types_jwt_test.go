package jwtKit

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	j := NewJWT([]byte("qwdqwdqwdq46456465"))
	method := jwt.SigningMethodHS256
	//method := jwt.SigningMethodHS384
	//method := jwt.SigningMethodHS512

	tokenString, err := j.New(method, map[string]interface{}{
		"a":   "b",
		"exp": jwt.NewNumericDate(time.Now().Add(-time.Hour)),
		//"exp": jwt.NewNumericDate(time.Now().Add(time.Hour)),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("tokenString", tokenString)

	mc, err := j.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return j.Key, nil
	})
	if err != nil {
		if IsTokenExpiredError(err) {
			panic("Token is expired.")
		}
		panic(err)
	}
	fmt.Println(mc)
}

func TestJWTComplexly(t *testing.T) {
	j := NewJWT([]byte("your_secret"), 42)
	method := jwt.SigningMethodHS256
	//method := jwt.SigningMethodHS384
	//method := jwt.SigningMethodHS512

	tokenString, err := j.NewComplexly(method, map[string]interface{}{
		"a": "b",
		"c": true,
		"d": 3.1415926,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("tokenString:", tokenString)

	mc, err := j.ParseComplexly(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return j.Key, nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(mc)
}
