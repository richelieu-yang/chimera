package gobKit

import (
	"bytes"
	"encoding/gob"
)

// Marshal 序列化.
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
