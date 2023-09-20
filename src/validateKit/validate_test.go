package validateKit

import (
	"fmt"
	"testing"
)

func TestValidateIPv4(t *testing.T) {
	fmt.Println(ValidateIPv4(""))           // Key: '' Error:Field validation for '' failed on the 'ipv4' tag
	fmt.Println(ValidateIPv4("127.0.0.1"))  // <nil>
	fmt.Println(ValidateIPv4("10.0.9.141")) // <nil>
	fmt.Println(ValidateIPv4("10.0.9141"))  // Key: '' Error:Field validation for '' failed on the 'ipv4' tag
}
