package main

import (
	"fmt"
	"github.com/richelieu42/chimera/v2/src/crypto/rsaKit"
)

func main() {
	//PKCS1
	pub := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAhCsm2ejKYy8a5FuedGpS\njiiLf/WM5o13MpHTBy9y4PuB88KjiO/TBDCSlSwwuHmVSXTERWjIF5drkbRa1c0E\nmAopYYt5387rmDb+FjumSh3BqtlrMQljkh7hG7JSXfpfoGNRUXi/WXAQoYW55MPz\nz7bEWL8a9DYLloRCzuysURZrIzhdsCdESmqwVakomQFqjCrL0agxwseF85Mr1U1N\nilFwmzYctn126uqYKYiDZRPbZj8yhOsCEyFWG0v+RRuMZJZuBgCM8KF8GTeeMNDX\nLp+53MEFAAjFdKAST+BluLbwWZE1c+hiQvHglpn5fmb5kbxSNNmMOhvkHcqP3mfK\nEQIDAQAB\n-----END PUBLIC KEY-----"
	pri := "-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEAhCsm2ejKYy8a5FuedGpSjiiLf/WM5o13MpHTBy9y4PuB88Kj\niO/TBDCSlSwwuHmVSXTERWjIF5drkbRa1c0EmAopYYt5387rmDb+FjumSh3Bqtlr\nMQljkh7hG7JSXfpfoGNRUXi/WXAQoYW55MPzz7bEWL8a9DYLloRCzuysURZrIzhd\nsCdESmqwVakomQFqjCrL0agxwseF85Mr1U1NilFwmzYctn126uqYKYiDZRPbZj8y\nhOsCEyFWG0v+RRuMZJZuBgCM8KF8GTeeMNDXLp+53MEFAAjFdKAST+BluLbwWZE1\nc+hiQvHglpn5fmb5kbxSNNmMOhvkHcqP3mfKEQIDAQABAoIBAAD2z95tb9sB/SUq\n7CyrGSGxduWKvvGwWN0PAmIiOOk+Wm4C1qatQa29VTdmem9BD0tJMVSUiWefJpUq\nMrKuL300r0U3c07UM4Sha5Bx1FJkdc0DT6Bsoive4utc3VQTS0roKI/ipKowSqXD\n/AV4vrS+X3+qerXK9mINYk+gOzhO8dEU3VtaUIaQaFMWtQaohbEXa4OAaXwajj5I\n1F6BLYreiH+5WAxcQkiKF34B5uqxpm7nmrrx9IM/zdXVIuuIbHQZgoog9TerAWFU\nN3qLKzMTcOncsMQCSCd2EKgwp644hxPittjbqVgi6fM0eY28pdADtEm7FdCOO5sd\nSXB7NqECgYEAiL9rypxOw5jh4K+oPUyMQJw/CiVuZNjtyhMrUOhWqZlkAhekw0/Z\nFWH4s3Ddz17o7faF2XUoVzhdOrbjFFbCE6yJJsAo7v+h822biZbKzsP9q5jjpqdi\nnB3jH0eGRHxZC0o7kHTJdcI/wDD9xVbbAfIjJM/huDBTkZpDMAXSgTECgYEA921w\n+JI7KnGC+vVoPd7Snvt5gYW1vcb8/bpHk4agENsINcf2AqRLNgiX6dKHEQPSkdZi\njxYnbhEd0vXYZU4vtdHgJ8WPI8f6/WWCJSEoNTM8BCERnH2jw8yhYADJ2YvGubKj\nVYD5Q4/aRp8D6vIwzmdlVxBcYGV150u6N7UonuECgYB7+L4qIuatIK62Wck2OW11\nLbFg62pXduqsphBA+0GF5A/sba7rejodzoH/e7U7SJn3EeVqrXHzEKBEmITpbXXp\nDCQSw2bYtVwHKjk11UTrincw2Byae2lEizvaoacx82FugM/bOGVjosTU4hVOkF1g\nODpwJx5FM1qvx6BdqwmaIQKBgAqLbnoZeU4IwmVBmiyw11cRLo93jFRHK7cAflZq\nMV4mh4YLPI/GQrJN0XOUlk8CU7IFafPhJNMLBbNc6NkAaQYdqKjpQX1r42VzMwW0\nEQubYF5ormB19km86c+2mOOdkl2NoGpfccAQXGQQGNGlerEBNQ9t0hIFoO23i3se\nTFnhAoGAL9ttoeap+GLo9arZDXUmrDVkbV0vApYgj6eGCGLao7MbQvsxmJbUfJYq\n1nn58ty/3TpuEYuvrfKhkCG9GoIHSrFmR5sv0WyYP3eVQ/RxUla6aTC42gWnwQA6\ne6AFD0CIrMULuaIXhVbTwTVZQhFm1hE1/atvOcQPzSEIp0aqcZY=\n-----END RSA PRIVATE KEY-----"

	data, err := rsaKit.PKCS1Encrypt([]byte("c强无敌群无多yy强无敌群"), []byte(pub))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	data, err = rsaKit.PKCS1Decrypt(data, []byte(pri), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
