package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func main() {
	key := []byte("my_secret_key")

	/*
		Create a new token object, specifying signing method and the claims
		you would like it to contain.
	*/
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	/*
		Sign and get the complete encoded token as a string using the secret
	*/
	tokenString, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("tokenString:", tokenString)
	fmt.Println("======")

	{
		// Parse takes the token string and a function for looking up the key. The latter is especially
		// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
		// head of the token to identify which key to use, but the parsed token (head and claims) is provided
		// to the callback, providing flexibility.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return key, nil
		})
		if err != nil {
			panic(err)
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			panic("not ok")
		}
		fmt.Println(claims["foo"])
		fmt.Println(claims["nbf"])
	}
}
