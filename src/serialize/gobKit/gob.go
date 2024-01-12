package gobKit

import (
	"bytes"
	"encoding/gob"
)

// Marshal 序列化.
/*
e.g.
	var m map[interface{}]interface{} = nil
	fmt.Println(gobKit.Marshal(m)) // [13 127 4 1 2 255 128 0 1 16 1 16 0 0 4 255 128 0 0] <nil>

	var obj interface{} = nil
	fmt.Println(gobKit.Marshal(obj)) // [] gob: cannot encode nil value
*/
func Marshal(obj any) ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	err := encoder.Encode(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Unmarshal 反序列化.
func Unmarshal(data []byte, ptr interface{}) error {
	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)

	return decoder.Decode(ptr)
}
