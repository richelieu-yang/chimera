package k8sYamlKit

import (
	"fmt"
	"testing"
)

func TestMarshalAndUnmarshal(t *testing.T) {
	type Person struct {
		Name string `json:"name"` // Affects YAML field names too.
		Age  int    `json:"age"`
	}

	// Marshal a Person struct to YAML.
	p := Person{"John", 30}
	y, err := Marshal(p)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(string(y))
	/* Output:
	age: 30
	name: John
	*/

	// Unmarshal the YAML back into a Person struct.
	var p2 Person
	err = Unmarshal(y, &p2)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println(p2)
	/* Output:
	{John 30}
	*/
}
