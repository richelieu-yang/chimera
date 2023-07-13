package rsaKit

import "testing"

func TestGenerateKeyFiles(t *testing.T) {
	options := []RsaOption{
		WithFormat(PKCS1),
		WithPassword(""),
	}

	priPath := "_pri.key"
	pubPath := "_pub.key"
	if err := GenerateKeyFiles(4096, priPath, pubPath, options...); err != nil {
		panic(err)
	}
}
