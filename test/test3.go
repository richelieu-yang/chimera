package main

import (
	"bufio"
	"github.com/sirupsen/logrus"
	"os"
)

var path = "a.txt"

func main() {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i++
		logrus.Infof("%d %s", i, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
