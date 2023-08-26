package main

import (
	"encoding/json"

	"github.com/bytedance/go-tagexpr/v2/binding"
	"github.com/cloudwego/hertz/pkg/app/server/render"
)

func main() {
	// Render
	render.ResetJSONMarshal(json.Marshal)

	// Binding
	binding.ResetJSONUnmarshaler(json.Unmarshal)
}
