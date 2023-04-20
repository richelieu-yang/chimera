package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "http://localhost:10000/api.do"
	method := "GET"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("1", "2")
	_ = writer.WriteField("2", "4")
	_ = writer.WriteField("3", "")
	_ = writer.WriteField("4", "")
	file, errFile5 := os.Open("/Users/richelieu/Documents/ino/images/可达鸭.jpg")
	defer file.Close()
	part5, errFile5 := writer.CreateFormFile("f1", filepath.Base("/Users/richelieu/Documents/ino/images/可达鸭.jpg"))
	_, errFile5 = io.Copy(part5, file)
	if errFile5 != nil {
		fmt.Println(errFile5)
		return
	}
	file, errFile6 := os.Open("/Users/richelieu/Documents/ino/images/盟主.png")
	defer file.Close()
	part6, errFile6 := writer.CreateFormFile("f2", filepath.Base("/Users/richelieu/Documents/ino/images/盟主.png"))
	_, errFile6 = io.Copy(part6, file)
	if errFile6 != nil {
		fmt.Println(errFile6)
		return
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
