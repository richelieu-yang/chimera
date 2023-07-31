package main

import (
	"fmt"
	"github.com/golang/snappy"
)

func main() {
	fmt.Println([]byte("aaa"))
	src1 := []byte("akakakakakakakakakakdddddddddcccccceeeeeeeegggggggggsssss")

	var dst1 []byte
	c := snappy.Encode(dst1, src1)
	fmt.Printf("src1 before compression len:%d\n", len(src1))
	fmt.Printf("src1 after compression len:%d\n", len(c))
}
