package rocketmq5Kit

import "github.com/richelieu-yang/chimera/v2/src/core/sliceKit"

func MixTags(tags ...string) string {
	return sliceKit.Join(tags, "||")
}

func GetTagString(tag *string) string {
	if tag == nil {
		return ""
	}
	return *tag
}
