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

	fmt.Println(re)                             // .+\.yozo\.com
	fmt.Println(re.MatchString("11yozo2com"))   // false
	fmt.Println(re.MatchString(".yozo.com"))    // false
	fmt.Println(re.MatchString("1.yozo.com"))   // true
	fmt.Println(re.MatchString("www.yozo.com")) // true
}
