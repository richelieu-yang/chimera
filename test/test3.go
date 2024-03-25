package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"path/filepath"
	"regexp"
)

func main() {

	filepath.Walk()
	filepath.WalkDir()

	re := regexp.MustCompile("\\[(\\d+)[vV]*(\\d*)\\]")

	s := []string{
		"[Sakurato] Mato Seihei no Slave [01][AVC-8bit 1080p AAC][CHS].mp4",
		"[Sakurato] Mato Seihei no Slave [02v2][AVC-8bit 1080p AAC][CHT].mp4",
		"[Sakurato] Mato Seihei no Slave [03][AVC-8bit 1080p AAC][CHT].mp4",
	}
	for i, str := range s {
		if !re.MatchString(str) {
			continue
		}

		fmt.Printf("--- %d ---\n", i)
		// string 类型
		fmt.Println(re.ReplaceAllString(str, "$1"))
		// []string 类型
		fmt.Println(re.FindAllString(str, -1))
		fmt.Printf("--- %d ---\n", i)
	}
}
