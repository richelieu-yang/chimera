package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
)

func main() {
	pub := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAla+RIorh09OO3Y/T8aM/\n5cumJtCO1oRZ3HBBJnXavNKUaVXgRCXNfsUeqoo8KSmTdaJnILQQc6BNEZePJURG\n5Wnls+pNUsHAve76cHAaEeZYSys/kd/RwO3t9WP8VqVjghzie0pA4Spr0ZEXNUg5\nq+XlcpQv9M+Cm869t+abG+eWc35KC0f59rqunbpMePJqF4Vyx50Z1YuRTVC7rV0Z\non6DwtexbHiJqm5UKYa7Rg3pPVPzP0z0nyXtV/xwk/QOff9WLswQYXEEO3Y/N9ZK\nkQAl3bPzag9u7OKKAuR3I+Xsm9OmpmFnqd3SVk+iyawdx855nTIMu3VVWzqAI2mM\nvwIDAQAB\n-----END PUBLIC KEY-----"
	pri := "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAla+RIorh09OO3Y/T8aM/5cumJtCO1oRZ3HBBJnXavNKUaVXg\nRCXNfsUeqoo8KSmTdaJnILQQc6BNEZePJURG5Wnls+pNUsHAve76cHAaEeZYSys/\nkd/RwO3t9WP8VqVjghzie0pA4Spr0ZEXNUg5q+XlcpQv9M+Cm869t+abG+eWc35K\nC0f59rqunbpMePJqF4Vyx50Z1YuRTVC7rV0Zon6DwtexbHiJqm5UKYa7Rg3pPVPz\nP0z0nyXtV/xwk/QOff9WLswQYXEEO3Y/N9ZKkQAl3bPzag9u7OKKAuR3I+Xsm9Om\npmFnqd3SVk+iyawdx855nTIMu3VVWzqAI2mMvwIDAQABAoIBABH40cIwo/XyRrWk\nw9vABWH/1jkpfUeVqHxmdfa6MRh5aBxQP+xZZZeSRKD76/E9RAPcYJQf9SxiKMld\n8G/sXLQax4ZB3CV9jI4Wwrb79AVbf2pht5uPdfhmGNYHuo+6An09c+voTxXkfqww\nxsqSRhUKNfGl3S7VjhGgUaMM6xP/G+ZSDoFny0dKOjDPhnSqzPM4+3rdAihpe0ZA\nqqhaYeu71W6iHI3pDXuRrc0OLK/XSyJWcvTcLyYa2VS8XFb6CMEoGr6JX/CtyBzH\n3/CzX2ldRHdTOU4XrRwzEqY0W+vqdi7Fw2bPBbQ4iUigum43DGpqVOSSNKl9xQya\no1C2niECgYEAuV4aadM5MlPymmaWSbPCQHLBqTjKVSkIQ0G7IqwGQNChpJvA5vkA\naXuRRzoEhDmuGfCg6H92h8qa1Eg0vxeiJXvsstCS/C0yEYVIda47ApE2EAHK62u5\nn6q3nls3EfUAPlxLdWRVkHuUJ9bhF4taEfQrBhIodBQXcVsws/He49kCgYEAzrjX\nx2FjUxXLerqnGo6p2kgF6sqEslJiLHT/LHtO6+GHz7HUS72o5zpoZ1tDLB3Knroz\nd1lW+lgcgBP/cYYAwdd389u2xnIigl0qfRtnRDOEqcLvJhlUCQsG+asMvN1Pljgz\n5thA1TA11hlW48VEZ5ysx1JTzItwPnRY1iBuTlcCgYAhZqx5vx7CS9AeNDgGZcat\n2D+/H0WOHMupWtG+iKLD0RQfVhVnSRzi5OojnHet37rYX7yOyFDyXPaTI8OlXzvP\nZ0oaWmzx1a786zhIRDHL36CPN9N5ojqbY6aceTEuIih+FCFgsMieNTZYC2cZ1/lN\n3OJXwLPbXrdurJDVC9Mu2QKBgQDGkXdz5XBrFrxmH3T0ZaVyL0y/w9jDEa+L1LZo\nG+orWnACaWbrejsMIWa4IUnczqEa5vEY65BI2OpHfO4aNgX0LOzkKWDgLjJH1wZB\nNllzgaklCEJ/7kRyi00f9dY+dLdGLKZuT7u3DKAqEB0OnSagADdXHxWvnSEsXMYY\nqP5L8QKBgGDWYqs4vgBc7EBD8Xj4kip9NzpcdklhfYAby9+Tu3LGqZqxnRCrdn49\nR9UAny1hgk7iRvs+jxIm3B986Yex7pvWPh0282YNQz9BPSG3NioN9GIlKejkxmpM\nFG33lm5emhh/Pt3DAY38r9NwFvQnA067sFFyTsjP8d3ULF2NeNuA\n-----END RSA PRIVATE KEY-----"

	data, err := rsaKit.Encrypt([]byte("cyy"), []byte(pub))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	data, err = rsaKit.Decrypt(data, []byte(pri), nil)
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
