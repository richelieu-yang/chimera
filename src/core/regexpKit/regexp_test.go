package regexpKit

import (
	"fmt"
	"testing"
)

func TestStringToRegexp(t *testing.T) {
	re, err := StringToRegexp("**.yozo.com")
	if err != nil {
		panic(err)
	}
	fmt.Println(re)
	fmt.Println(re.MatchString("11yozo2com"))
	fmt.Println(re.MatchString(".yozo.com"))
	fmt.Println(re.MatchString("1.yozo.com"))
	fmt.Println(re.MatchString("www.yozo.com"))
}
