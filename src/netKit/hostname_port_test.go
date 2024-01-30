package netKit

import (
	"fmt"
	"testing"
)

func TestSplitHostnamePort(t *testing.T) {
	fmt.Println(SplitHostnamePort("localhost:8080")) // "localhost" "8080" <nil>

	fmt.Println(SplitHostnamePort("127.0.0.1")) // "" "" address 127.0.0.1: missing port in address
	fmt.Println(SplitHostnamePort("localhost")) // "" "" address localhost: missing port in address
}
