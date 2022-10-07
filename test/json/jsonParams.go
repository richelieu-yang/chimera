package main

import (
	"gitee.com/richelieu042/go-scales/src/core/errorKit"
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	jsoniter "github.com/json-iterator/go"
)

type (
	JsonParams struct {
		Method int                    `json:"method"`
		Params map[string]interface{} `json:"params"`
	}
)

// NewFromJson
/**
 * 方法名参考：decimal库
 * e.g.
 *
 */
func NewFromJson(json string) (*JsonParams, error) {
	// 空字符串要特殊处理
	// 否则传到下面去的话会报错：readObjectStart: expect { or n, but found  , error found in #0 byte of ...||..., bigger context ...||...
	if strKit.IsEmpty(json) {
		return nil, nil
	}

	jsonParams := &JsonParams{}
	// byte <=> uint8
	if err := jsoniter.Unmarshal([]byte(json), jsonParams); err != nil {
		return nil, errorKit.Wrap(err, "NewFromJson() fails with json(%s)", json)
		//return nil, errors.Simple(strKit.Format("反序列化jsonParams字符串失败，error: [%v]", err.Error()))
	}
	return jsonParams, nil
}
