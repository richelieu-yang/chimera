package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"github.com/tidwall/gjson"
)

func main() {
	{
		json := `{"status":"1","info":"OK","infocode":"10000","province":[],"city":[],"adcode":[],"rectangle":[]}`
		f := jsonKit.GetField([]byte(json), "province")
		fmt.Println(f)
		fmt.Println(f.Type)
		fmt.Println("IsString", f.Type == gjson.String)
		fmt.Println("IsArray", f.IsArray())
	}
	fmt.Println("------")
	{
		json := `{"status":"1","info":"OK","infocode":"10000","province":"局域网","city":[],"adcode":[],"rectangle":[]}`
		f := jsonKit.GetField([]byte(json), "province")
		fmt.Println(f)
		fmt.Println(f.Type)
		fmt.Println("IsString", f.Type == gjson.String)
		fmt.Println("IsArray", f.IsArray())
	}
	/*
		[]
		JSON
		IsString false
		IsArray true
		------
		局域网
		String
		IsString true
		IsArray false
	*/
}
