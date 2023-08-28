package reqKit

import (
	"errors"
	"fmt"
	"io"
	"testing"
)

func TestNewClient(t *testing.T) {
	{
		/*

			PS:
			(1) 通过判断 resp.Err 来判断是否发生 error.
			(2) 任何情况下返回的 “resp” 永远不会是 nil，可以放心的直接判断。
		*/
		url := "https://www.baidu.com/"

		client := NewClient()
		resp := client.Get(url).Do()
		if resp.Err != nil {
			panic(resp.Err)
		}
		defer resp.Body.Close()

		if !resp.IsSuccessState() {
			panic(errors.New(fmt.Sprintf("error state: %d", resp.StatusCode)))
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	}

}
