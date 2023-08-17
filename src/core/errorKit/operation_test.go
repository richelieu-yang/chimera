package errorKit

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestIs(t *testing.T) {
	err := redis.Nil
	err1 := Wrap(err, "1")
	err2 := Wrap(err1, "2")

	fmt.Printf("%+v\n", err2)

	fmt.Println(Is(err2, err)) // true
	fmt.Println(Is(err1, err)) // true
	fmt.Println(Is(err, err))  // true

	fmt.Println(Is(err2, err1)) // true
	fmt.Println(Is(err1, err2)) // false
}
