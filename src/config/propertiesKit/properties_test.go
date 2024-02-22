package propertiesKit

import (
	"fmt"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	text := `NaMe=Richelieu
Age=3
#c=cyy
`
	//var obj interface{}
	var obj map[string]string
	if err := Unmarshal([]byte(text), &obj); err != nil {
		panic(err)
	}
	fmt.Println(obj)
}
