package baiduOcrKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/http/httpClientKit"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"sync"
	"time"
)

var (
	// clientId 即 apiKey
	clientId = ""
	// clientSecret 即 secretKey
	clientSecret = ""
)

var lock = new(sync.Mutex)
var token *accessToken = nil

func SetApiKeyAndSecretKey(apiKey, secretKey string) error {
	lock.Lock()
	defer lock.Unlock()

	apiKey = strKit.Trim(apiKey)
	secretKey = strKit.Trim(secretKey)
	if strKit.IsEmpty(apiKey) {
		return errorKit.Simple("apiKey is empty")
	}
	if strKit.IsEmpty(secretKey) {
		return errorKit.Simple("secretKey is empty")
	}
	token = nil
	clientId = apiKey
	clientSecret = secretKey
	return nil
}

func getAccessToken() (*accessToken, error) {
	lock.Lock()
	defer lock.Unlock()

	// (1) token存在
	if token != nil {
		if !token.isExpired() {
			return token, nil
		}
		// 去除已过期的token
		token = nil
	}
	// (2) token存在，重新生成
	var err error
	token, err = newAccessToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

// newAccessToken 向百度获取 Access Token
/*
参考: https://ai.baidu.com/ai-doc/REFERENCE/Ck3dwjhhu

@return 必定一个为nil，一个非nil
*/
func newAccessToken() (*accessToken, error) {
	if strKit.HasEmpty(clientId, clientSecret) {
		return nil, errorKit.Simple("SetApiKeyAndSecretKey() should be invoked firstly")
	}

	url := strKit.Format("https://aip.baidubce.com/oauth/2.0/token?grant_type=%s&client_id=%s&client_secret=%s",
		grantType,
		clientId,
		clientSecret)
	resp, err := httpClientKit.Get(url, nil)
	if err != nil {
		return nil, err
	}
	m, err := jsonKit.UnmarshalToMap(resp)
	if err != nil {
		return nil, err
	}

	token, err := parseMapToAccessToken(m)
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errorKit.Simple("failure response(%s)", string(resp))
	}
	// 对token进行简单验证
	// 正常情况下，token的有效期为30天，此处判断是为了防止特殊情况（86400秒 == 1天）
	if token.ExpiresIn <= 86400 {
		return nil, errorKit.Simple("token.ExpiresIn(%d) is invalid", 86400)
	}
	// 提前1h(3600s)认为token失效，以防特殊情况
	token.ExpireTime = time.Now().Add(time.Second * time.Duration(token.ExpiresIn-3600))
	return token, nil
}
