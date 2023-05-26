package rsaKit

type (
	KeyFormat uint8
)

const (
	PKCS1 KeyFormat = iota
	PKCS8
)
