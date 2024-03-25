package main

import (
	"fmt"
	"html"
)

func main() {
	htmlStr := "<html>Hello.</html>"
	escaped := html.EscapeString(htmlStr)
	fmt.Println(escaped)
	htmlStr1 := html.UnescapeString(escaped)
	fmt.Println(htmlStr1)

	//fmt.Println(base64Kit.DecodeStringToString("cG9uZw=="))

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
