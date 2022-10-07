package script

import (
	"fmt"
	"gitee.com/richelieu042/go-scales/src/database/redisKit/test"
	"testing"
)

func TestExists(t *testing.T) {
	fmt.Println("测试 Exists ---------------------------------------")

	// client
	client, err := test.NewSingleNodeClient()
	if err != nil {
		panic(err)
	}

	/*
		redis.call("set", "1", "a")
		return redis.call("get", "1")
	*/
	script := "redis.call(\"set\", \"1\", \"a\") return redis.call(\"get\", \"1\")"
	sha, err := client.ScriptLoad(script)
	if err != nil {
		panic(err)
	}
	fmt.Printf("sha: [%s].\n", sha)
	value, err := client.EvalSha(sha, nil)
	if str, ok := value.(string); ok {
		if str != "a" {
			panic("get和set的值不相等！！！")
		}
		fmt.Println("测试 Exists ---------------------------------------")
	} else {
		panic("执行返回值的类型不是string！！！")
	}
}
