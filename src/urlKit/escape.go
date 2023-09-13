package urlKit

import "net/url"

// EncodeURIComponent 编码.
/*
e.g.
("") => ""
*/
var EncodeURIComponent func(s string) string = url.QueryEscape

// DecodeURIComponent 解码.
/*
e.g.
("") => "", nil
*/
var DecodeURIComponent func(s string) (string, error) = url.QueryUnescape
