package rsaKit

//// parsePublicKey 解析公钥（PKCS1 || PKCS8）
//func parsePublicKey(s []byte) (*rsa.PublicKey, error) {
//	block, _ := pem.Decode(s)
//	if block == nil {
//		return nil, errorKit.Simple("public key error")
//	}
//
//	keyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	return keyInterface.(*rsa.PublicKey), nil
//}

//// parsePKCS1PrivateKey 解析私钥（PKCS1）
//func parsePKCS1PrivateKey(s []byte) (*rsa.PrivateKey, error) {
//	block, _ := pem.Decode(s)
//	if block == nil {
//		return nil, errorKit.Simple("private key error(%s)", "PKCS1")
//	}
//
//	return x509.ParsePKCS1PrivateKey(block.Bytes)
//}
//
//// parsePKCS8PrivateKey 解析私钥（PKCS8）
//func parsePKCS8PrivateKey(s []byte) (*rsa.PrivateKey, error) {
//	block, _ := pem.Decode(s)
//	if block == nil {
//		return nil, errorKit.Simple("private key error(%s)", "PKCS8")
//	}
//
//	keyInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
//	if err != nil {
//		return nil, err
//	}
//	return keyInterface.(*rsa.PrivateKey), nil
//}
