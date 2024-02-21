package versionKit

import (
	"fmt"
	"testing"
)

func TestNewVersion(t *testing.T) {
	v, err := NewVersion("1.0.0-alpha.1+build.123")
	if err != nil {
		panic(err)
	}
	/*
		v.pre == "alpha.1"
		v.metadata == "build.123"
	*/
	fmt.Println(v)
}

func TestNewVersion1(t *testing.T) {
	{
		v, err := NewVersion("1.3.10+meta")
		if err != nil {
			panic(err)
		}
		fmt.Println(v) // 1.3.10+meta
	}

	// Richelieu: 会忽略最前面的'v'
	{
		v, err := NewVersion("v1.3.10+meta")
		if err != nil {
			panic(err)
		}
		fmt.Println(v) // 1.3.10+meta
	}
}

func TestCheckConstraint(t *testing.T) {
	fmt.Println(CheckConstraint("0.9", ">= 1.0, < 1.4")) // false <nil>

	fmt.Println(CheckConstraint("1.2", ">= 1.0, <= 1.2")) // true <nil>

	fmt.Println(CheckConstraint("v1.2", ">= 1.0, <= 1.2")) // true <nil>
	fmt.Println(CheckConstraint("v1.2", ">= 1.0,<= 1.2"))  // true <nil>
}

func TestSort(t *testing.T) {
	s := []string{"3.4", "4.0", "3.3", "4.0+metadata", "1.0.0-alpha+001"}
	fmt.Println(Sort(s))
}
