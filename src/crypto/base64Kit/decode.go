package base64Kit

func Decode(src []byte, options ...Base64Option) ([]byte, error) {
	opts := loadOptions(options...)
	return opts.Decode(src)
}

func DecodeString(s string, options ...Base64Option) ([]byte, error) {
	opts := loadOptions(options...)
	return opts.DecodeString(s)
}

// DecodeStringToString (拓展) base64 string => string
func DecodeStringToString(s string, options ...Base64Option) (string, error) {
	data, err := DecodeString(s, options...)
	return string(data), err
}

//// Decode []byte => []byte
///*
//参考: gbase64.Decode()
//*/
//func Decode(data []byte, options ...Base64Option) ([]byte, error) {
//	opts := loadOptions(options...)
//
//	rst := make([]byte, opts.encoding.DecodedLen(len(data)))
//	n, err := opts.encoding.Decode(rst, data)
//	if err != nil {
//		return nil, err
//	}
//	return rst[:n], nil
//}
//
//// DecodeToString []byte => string
//func DecodeToString(data []byte, options ...Base64Option) (string, error) {
//	rst, err := Decode(data, options...)
//	return string(rst), err
//}
//
//// DecodeString string => []byte
//func DecodeString(str string, options ...Base64Option) ([]byte, error) {
//	return Decode([]byte(str), options...)
//}
//
//// DecodeStringToString string => string
//func DecodeStringToString(str string, options ...Base64Option) (string, error) {
//	rst, err := DecodeString(str, options...)
//	return string(rst), err
//}
