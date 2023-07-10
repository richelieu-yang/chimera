package main

import (
	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	sonic.ConfigDefault
	jsoniter.ConfigDefault
}
