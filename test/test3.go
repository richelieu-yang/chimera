package main

import (
	"fmt"
	"github.com/apache/rocketmq-clients/golang/v5/pkg/utils"
)

func main() {
	fmt.Println(string(utils.GetMacAddress()))
}
