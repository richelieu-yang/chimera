package charsetKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"testing"
)

func TestIsGBK(t *testing.T) {
	data, err := fileKit.ReadFile("_gbk.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(IsGBK(data))

	//type args struct {
	//	data []byte
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := IsGBK(tt.args.data); got != tt.want {
	//			t.Errorf("IsGBK() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
