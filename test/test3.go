package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
)

func main() {
	fmt.Println(base64Kit.DecodeStringToString("cG9uZw=="))

	//fmt.Println(base64Kit.DecodeStringToString(""))

	//password := []byte("MyDarkSecret")
	//
	//// Hashing the password with the default cost of 10
	//hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(hashedPassword))
	//
	//// Comparing the password with the hash
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	//fmt.Println(err) // nil means it is a match
}
