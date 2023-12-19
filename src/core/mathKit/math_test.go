package mathKit

import (
	"fmt"
	"math"
	"testing"
)

func TestSin(t *testing.T) {
	fmt.Println(Sin(0))           // -0
	fmt.Println(Sin(45))          // 0.851
	fmt.Println(Sin(90))          // 0.894
	fmt.Println(Sin(math.Pi))     // 0
	fmt.Println(Sin(math.Pi / 2)) // 1
}

func TestCos(t *testing.T) {
	fmt.Println(Cos(0))           // 1
	fmt.Println(Cos(45))          // 0.525
	fmt.Println(Cos(90))          // -0.447
	fmt.Println(Cos(math.Pi))     // -1
	fmt.Println(Cos(math.Pi / 2)) // -0
}
