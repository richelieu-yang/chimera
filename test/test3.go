package main

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/mapping"
)

func main() {
	type inner struct {
		Value1 int `key:"value1,range=(100:]"`
		//Value2  int8   `key:"value2,range=[1:5]"`
		//Value3  int16  `key:"value3,range=[1:5]"`
		//Value4  int32  `key:"value4,range=[1:5]"`
		//Value5  int64  `key:"value5,range=[1:5]"`
		//Value6  uint   `key:"value6,range=[:5]"`
		//Value8  uint8  `key:"value8,range=[1:5],string"`
		//Value9  uint16 `key:"value9,range=[1:5],string"`
		//Value10 uint32 `key:"value10,range=[1:5],string"`
		//Value11 uint64 `key:"value11,range=[1:5],string"`
	}
	m := map[string]any{
		"value1": 100,
		//"value2":  int8(1),
		//"value3":  int16(2),
		//"value4":  int32(4),
		//"value5":  int64(5),
		//"value6":  uint(0),
		//"value8":  "1",
		//"value9":  "2",
		//"value10": "4",
		//"value11": "5",
	}

	var in inner
	if err := mapping.UnmarshalKey(m, &in); err != nil {
		panic(err)
	}
	fmt.Println(666)
}
