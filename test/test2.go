package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const API_KEY = "vdRY6PcvUuaXcCUPWRlnXrHM"
const SECRET_KEY = "zVMxRwFb3ZUeG6SGuggh9FOwTUfH6HQO"

func main() {
	url := "https://aip.baidubce.com/rpc/2.0/tts/v1/create?access_token=" + GetAccessToken()
	payload := strings.NewReader(`{"format":"mp3-16k","voice":0,"lang":"zh","speed":5,"pitch":5,"volume":5,"enable_subtitle":0,"text":"测试"}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

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

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}
