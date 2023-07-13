package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	MaxAge = time.Duration(60*60*24) * time.Second

	SECRETKEY = "243223ffslsfsldfl412fdsfsdf" //私钥
)

// CustomClaims 自定义Claims
type CustomClaims struct {
	jwt.RegisteredClaims

	UserName string
}

func main() {
	/* (1) 生成token */
	//claims := &jwt.MapClaims{
	//	"id":   11,
	//	"name": "jerry",
	//	"exp":  time.Now().Add(MaxAge).Unix(), // 过期时间，必须设置,
	//}
	claims := &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Second * 3),
			},
		},
		UserName: "测试",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		panic(err)
	}
	println("token:", tokenString)

	/* (2) 解析token */
	rst, err := ParseToken(tokenString)
	if err != nil {
		panic(err)
	}
	println("claims:", rst)
}

// ParseToken 解析token
func ParseToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return token.Claims, nil
}
