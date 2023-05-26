package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
)

func main() {
	pub := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnqt24Cq7wnbdzUZqh9FP\nKjLzsoXlRqfc5d8ivJNnSBYl3CfLRhipRsCMxzaOY9QR8wpP4cYgG9tSF//s2SLI\nsziAs9VSxZPjjj7i9UTLbE9IhuwT2X3SMqWDVQSRgUzgCT6C61MSxIxnT0ffn6+g\nlAcNXOOdDFE576F3KQTz65gdOWuRa2Sxigxy6N1Acl8Kmp6QqLH5OGHjjiKWISX+\nUWU5ZsiRdchpkhUGSxxBs0YBQvT9Aw+tte0BNOmjbCH3/xwT/jemq4peVKmW2ce2\n40O/KMlWkWt0V+PF/Q4NlLiOjMiVCOa0LI095/U+lABO9guXm1jCdmwRKnSoxYm/\nowIDAQAB\n-----END PUBLIC KEY-----\n"
	pri := "-----BEGIN PRIVATE KEY-----\nMIIEvwIBADANBgkqhkiG9w0BAQEFAASCBKkwggSlAgEAAoIBAQCeq3bgKrvCdt3N\nRmqH0U8qMvOyheVGp9zl3yK8k2dIFiXcJ8tGGKlGwIzHNo5j1BHzCk/hxiAb21IX\n/+zZIsizOICz1VLFk+OOPuL1RMtsT0iG7BPZfdIypYNVBJGBTOAJPoLrUxLEjGdP\nR9+fr6CUBw1c450MUTnvoXcpBPPrmB05a5FrZLGKDHLo3UByXwqanpCosfk4YeOO\nIpYhJf5RZTlmyJF1yGmSFQZLHEGzRgFC9P0DD6217QE06aNsIff/HBP+N6aril5U\nqZbZx7bjQ78oyVaRa3RX48X9Dg2UuI6MyJUI5rQsjT3n9T6UAE72C5ebWMJ2bBEq\ndKjFib+jAgMBAAECggEBAJJtyLINneNUEVN8tZmL15QpG29YVAQFrD1T1FnVas5C\n6kOnUdfpsGaEEidSuMt2De6OSLQiRMWQfM17ONgWWsPS1CuiJdYCqlE+xfL1vwnA\njT0+pv7jN9XXm0Ediy6ZhHPjhayoqLxMmpa2zLAww294pqNfxzJzz5TPHQn08Sos\nbmVpQAx9lhzXECesVp3YuPqmzLR/ni6qJyIKjVCZThLUNvra5044HHyxADftSb7x\nkQk/eARXRdHkrecvqAwc56rH63LZr1obVAcGXpfVlJlu41GYb03RNnDke5PUAXMC\n65UuXdUqNE7Px3vyX+ktrVTEyaT3uieJrCcF3UgZTNkCgYEAyZ4kzKAgS19NwAnJ\nRmvxh89uzZ61shu0BA2gmb837+FEqXsfNN+qvwePSpVlZ+JapolmX1E9EnSeGwBM\nQUbB4rRhXUKNmdmyMBmFvoWlxLD9nLaFh8DouJLtOwsUttUlDkQlSJj3m0nua+B8\nJKsh8CweCO5YpolSs2Xku5Uit+UCgYEAyXe7P1UwKgBG81DeZ+xMNEq1A+rGsVXl\nWAlneHOplbhICZgs9Onb2Bt5e+1TdYu+EYZTgCE5CKHrEjFS1eHzyiMGGQskSU5t\niUnQL8jXTvX/bYMIIpz0K27Fr8GIN8XrYxgpHLgpRAMzCoMPn0BJ/myp0912snQq\nlzC89iZhkOcCgYBGsH3yiMoJ4zZdeJDrcflTRrVwvo6yDeiF60k93r98CZ2LVrHQ\nVgOFyt7ApbTSQZjK/y23icJB8EVm7inOUUK1e80PZ8BCCwPgvXIZJL7EQdIJ2izj\nNg3ieRFYPEBdAkplwBraY1edSMShaincqWvSy4UUWI0YhtlKMRWLd3SeyQKBgQC2\nvYC9qJKgdlHk38RsdTyF7gGVBP6m2efGW/lbm2YGZPoKodqdaAY+VmVyEEm0hEWp\n9bKFtRzbhKsJfG3LjUEpZ21fh+ipCZd3gNlyjnUlrSK7/a+CXeVoA/kSO7RICqFA\n6ChaYX0ksNY6QbkI1TKIsZ21i9rfcVk0fEPZ7eRvnwKBgQC96FbkxUTQTibgbSTl\nAndnOA2uSV1Z59jDQX0xJVxBGJDEXinr16iWxoDLjtvsu8sW0WrJnz5rbkZs4es3\nuHSGrUsmgHZawWc97H4CxPQX/k2DVQxjLtFrJJ6YyoiCUiXa5VO+hBcwHPv5lmTg\nQY435L0OJjTpef1T/Ar3SYEgGQ==\n-----END PRIVATE KEY-----\n"

	data, err := rsaKit.PKCS8Encrypt([]byte("c强无敌群无多yy强无敌群"), []byte(pub))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	data, err = rsaKit.PKCS8Decrypt(data, []byte(pri), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

//// RSA公钥私钥产生
//func GenRsaKey(bits int) error {
//	// 生成私钥文件
//	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
//	if err != nil {
//		return err
//	}
//	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
//	block := &pem.Block{
//		Type:  "RSA PRIVATE KEY",
//		Bytes: derStream,
//	}
//	file, err := os.Create("private.pem")
//	if err != nil {
//		return err
//	}
//	err = pem.Encode(file, block)
//	if err != nil {
//		return err
//	}
//	// 生成公钥文件
//	publicKey := &privateKey.PublicKey
//	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
//	if err != nil {
//		return err
//	}
//	block = &pem.Block{
//		Type:  "PUBLIC KEY",
//		Bytes: derPkix,
//	}
//	file, err = os.Create("public.pem")
//	if err != nil {
//		return err
//	}
//	err = pem.Encode(file, block)
//	if err != nil {
//		return err
//	}
//	return nil
//}
