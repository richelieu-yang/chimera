package pathKit

import (
	"fmt"
	"testing"
)

func TestChangeWorkingDir(t *testing.T) {
	wd, err := GetWorkingDir()
	if err != nil {
		panic(err)
	}
	fmt.Println("wd: ", wd)

	tmp := Join(wd, "1", "a")
	fmt.Println("tmp: ", tmp)
	if err := ChangeWorkingDir(tmp); err != nil {
		panic(err)
	}

	wd1, err := GetWorkingDir()
	if err != nil {
		panic(err)
	}
	fmt.Println("wd1: ", wd1)

	if tmp == wd1 {
		fmt.Println("equal")
	} else {
		panic("not equal")
	}

}
