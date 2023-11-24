package main

import "github.com/pkg/browser"

func main() {
	browser.Stderr

	if err := browser.OpenURL("https://www.baidu.com"); err != nil {
		panic(err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	//defer cancel()
	//
	//if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
	//	fmt.Println("got error:", err) // got error: signal: killed
	//}
}
