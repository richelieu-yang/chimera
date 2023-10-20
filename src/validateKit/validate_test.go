package validateKit

import (
	"fmt"
	"testing"
)

func TestValidateIPv4(t *testing.T) {
	fmt.Println(IPv4(""))           // Key: '' Error:Var validation for '' failed on the 'ipv4' tag
	fmt.Println(IPv4("127.0.0.1"))  // <nil>
	fmt.Println(IPv4("10.0.9.141")) // <nil>
	fmt.Println(IPv4("10.0.9141"))  // Key: '' Error:Var validation for '' failed on the 'ipv4' tag
}
