package main

import jose "github.com/dvsekhvalnov/jose2go"

func init() {
	/*
		TODO: 看后续 github.com/apache/pulsar-client-go v0.11.1 会不会更新.
		Richelieu: 为了引用最新的 dvsekhvalnov/jose2go 库（v1.6.0），v1.5.0有脆弱性漏洞，
	*/
	_ = jose.A128CBC_HS256
}
