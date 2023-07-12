package main

import "fmt"

func main() {
	var r rune = 61
	fmt.Println(string(r))

	//a := base64.StdEncoding
	//a.WithPadding(base64.NoPadding)
	//a.WithPadding(base64.NoPadding)

	//input := []byte("\x00\x00\x3e\x00\x00\x3f\x00")
	//fmt.Println(base64.StdEncoding.EncodeToString(input)) // AAA+AAA/==
	//fmt.Println(base64.URLEncoding.EncodeToString(input)) // AAA-AAA_==
	//fmt.Println(base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(input)）// AAA+AAA/
	//fmt.Println(base64.URLEncoding.With(base64.NoPadding).EncodeToString(input)）// AAA-AAA_

}
