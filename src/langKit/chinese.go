package langKit

import (
	"github.com/liuzl/gocc"
)

// T2S 繁体中文 => 简体中文
func T2S(in string) (string, error) {
	return convert("t2s", in)
}

// S2T 简体中文 => 繁体中文
func S2T(in string) (string, error) {
	return convert("s2t", in)
}

// S2HK 简体中文 => 繁体中文（香港）
func S2HK(in string) (string, error) {
	return convert("s2hk", in)
}

// S2TW 简体中文 => 繁体中文（台湾）
func S2TW(in string) (string, error) {
	return convert("s2tw", in)
}

// convert
/*
@param conversion s2t, t2s, s2tw, tw2s, s2hk, hk2s, s2twp, tw2sp, t2tw, t2hk
*/
func convert(conversion, in string) (string, error) {
	cc, err := gocc.New(conversion)
	if err != nil {
		return "", err
	}
	return cc.Convert(in)
}
