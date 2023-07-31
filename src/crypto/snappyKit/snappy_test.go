package snappyKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/bytesKit"
	"testing"
)

func TestEncode(t *testing.T) {
	//src := []byte("11111111")
	src := []byte("akakakakakakakakakakdddddddddcccccceeeeeeeegggggggggsssss")

	dst := Encode(src)
	fmt.Println("dst", string(dst), len(dst))

	src1, err := Decode(dst)
	if err != nil {
		panic(err)
	}
	fmt.Println("src1", string(src1), len(src1))

	if !bytesKit.Equals(src, src1) {
		panic("not equal")
	}
	fmt.Println("equal")
}
