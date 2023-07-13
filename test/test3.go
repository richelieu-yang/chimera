package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/crypto/rsaKit"
	"time"
)

const (
	MaxAge = timeKit.Day
)

// CustomClaims 自定义Claims
type CustomClaims struct {
	jwt.RegisteredClaims

	UserName string
}

var priPem []byte
var pubPem []byte

func main() {
	// 私钥的密码
	password := "dqwdqwd强无敌群多"
	var err error
	priPem, err = fileKit.ReadFile("_pri.pem")
	if err != nil {
		panic(err)
	}
	pubPem, err = fileKit.ReadFile("_pub.pem")
	if err != nil {
		panic(err)
	}

	privateKey, err := rsaKit.ParsePrivateKeyFromPem(priPem, password)
	if err != nil {
		panic(err)
	}
	publicKey, err := rsaKit.ParsePublicKeyFromPem(pubPem)
	if err != nil {
		panic(err)
	}

	/* (1) 生成token */
	//claims := &jwt.MapClaims{
	//	"id":   11,
	//	"name": "jerry",
	//	"exp":  time.Now().Add(MaxAge).Unix(), // 过期时间，必须设置,
	//}
	claims := &CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(MaxAge),
			},
		},
		UserName: "测试",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("jwt:\n%s\n", tokenString)

	/* (2) 解析token */
	rst, err := ParseToken(publicKey, tokenString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("claims:\n%v\n", rst)
}

// ParseToken 解析token
func ParseToken(key interface{}, tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		//return []byte(SECRETKEY), nil
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	return token.Claims, nil
}
