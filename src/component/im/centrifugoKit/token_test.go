package centrifugoKit

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func TestGenToken(t *testing.T) {
	secret := []byte("6bxFK7Oos3y3VwqJlr1gmoUrW-Yp3TP1Dol6ptDg4KI_zgra2Wk4JHjdmnOGXdpVppFZCndQ9yVEDSV88Sh_wA")
	token, err := GenerateToken(secret, jwt.SigningMethodHS256, "user", time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}

func TestGenerateSubToken(t *testing.T) {
	secret := []byte("6bxFK7Oos3y3VwqJlr1gmoUrW-Yp3TP1Dol6ptDg4KI_zgra2Wk4JHjdmnOGXdpVppFZCndQ9yVEDSV88Sh_wA")
	token, err := GenerateSubToken(secret, jwt.SigningMethodHS256, "user", "news", time.Hour)
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}
