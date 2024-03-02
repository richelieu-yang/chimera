package jwtKit

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"testing"
)

func TestComplexly(t *testing.T) {
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
