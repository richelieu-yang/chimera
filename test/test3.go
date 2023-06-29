package main

import (
	"gorm.io/gorm"
)

func main() {
	type User struct {
		gorm.Model

		Name string `json:"name"`
		Age  uint   `json:"age"`
	}

	//structs.Map()
}
