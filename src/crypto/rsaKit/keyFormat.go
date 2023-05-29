package rsaKit

import "fmt"

type (
	KeyFormat uint8
)

func (format KeyFormat) String() string {
	return fmt.Sprintf("PKCS#%d", format)
}

const (
	PKCS1 KeyFormat = 1
	PKCS8 KeyFormat = 8
)
