package jsonKit

import (
	"fmt"
	"testing"
)

func TestGetStringField(t *testing.T) {
	jsonStr := `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	fmt.Println(GetStringField([]byte(jsonStr), "name.last")) // Prichard
}

func TestGetStringFieldFromJson(t *testing.T) {
	jsonStr := `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	fmt.Println(GetStringFieldFromString(jsonStr, "name.last")) // Prichard
}
