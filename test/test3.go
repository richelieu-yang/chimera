package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/token/jwtKit"
)

func main() {
	tokenString := "uoZxrWsyEyZYKpY1DyYiYdH5sSY6YafNLSZ9.uoZbuXQyEzU3CTa3CTG0ETwiYcPfrWLZPSY6Ydx4uSYiYdLpPNZZPSY6Ydx4uSZ9.GTurBggP8vAPTp_rZ-Qtd56vWrt7Eo_gGlNQ8va7F9a"

	j := jwtKit.NewJWT([]byte("my_secret"), 42)
	//method := jwt.SigningMethodHS256

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
	fmt.Println(mc["fileId"]) // xxx
	fmt.Println(mc["userId"]) // xxx
}
