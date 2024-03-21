package bcryptKit

import (
	"fmt"
	"testing"
)

func TestHashPassword(t *testing.T) {
	fmt.Println(HashPassword([]byte("123")))
	fmt.Println(HashPassword([]byte("123")))
}
