package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	fmt.Println(jwt.GetAlgorithms())
}
