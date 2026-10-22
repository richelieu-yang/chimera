package centrifugoKit

import (
	"cmp"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu042/chimera/v3/src/core/error/errKit"
)

// DefaultTokenTTL 是 centrifugo gentoken 和 gensubtoken 默认的有效期（7 天）。
const DefaultTokenTTL time.Duration = 7 * 24 * time.Hour

// GenerateToken 使用指定算法生成 Centrifugo 连接 JWT。
//
// 参数：
//   - key：HMAC 密钥，类型为 []byte，对应 client.token.hmac_secret_key，不能为空。
//     密钥字节原样参与签名，不做 Base64 解码或去除空白。
//   - method：签名算法，传入 nil 时默认使用 jwt.SigningMethodHS256。
//     非 nil 时只能传入 jwt.SigningMethodHS256、jwt.SigningMethodHS384、jwt.SigningMethodHS512 三者之一。
//     推荐使用 jwt.SigningMethodHS256，
//     与 centrifugo gentoken 使用的算法一致；所选算法需与服务端的验签配置匹配。
//     自行构造的签名方法实例会返回错误。
//   - user：用户 ID，原样写入 sub；空字符串表示匿名用户，不写入 sub。
//   - ttl：从当前时间开始计算的有效期，类型为 time.Duration，例如 time.Hour 或 30*time.Minute。
//     ttl > 0 时设置 exp；ttl <= 0（包括 0 和任意负数）时不设置 exp，即没有过期时间，也不会因此报错。
//     不会自动使用默认有效期；如需 7 天，请显式传入 DefaultTokenTTL。
//
// JWT 的时间戳精度为秒，过期时间的小数秒部分会被截断。
func GenerateToken(key []byte, method *jwt.SigningMethodHMAC, user string, ttl time.Duration) (string, error) {
	return generateToken(key, method, user, "", ttl)
}

// GenerateSubToken 使用指定算法生成 Centrifugo 频道订阅 JWT。
//
// 参数：
//   - key：HMAC 密钥，类型为 []byte，不能为空，需由调用方选择并传入。
//     启用独立订阅 token 配置时使用 client.subscription_token.token.hmac_secret_key，
//     否则使用 client.token.hmac_secret_key；密钥字节原样参与签名，不做 Base64 解码或去除空白。
//   - method：签名算法，传入 nil 时默认使用 jwt.SigningMethodHS256。
//     非 nil 时只能传入 jwt.SigningMethodHS256、jwt.SigningMethodHS384、jwt.SigningMethodHS512 三者之一。
//     推荐使用 jwt.SigningMethodHS256，
//     与 centrifugo gensubtoken 使用的算法一致；所选算法需与服务端的订阅 token 验签配置匹配。
//     自行构造的签名方法实例会返回错误。
//   - user：用户 ID，原样写入 sub；空字符串表示匿名用户，不写入 sub。
//   - channel：订阅频道名，原样写入 channel，必须与客户端订阅的频道名完全一致，不能为空。
//   - ttl：从当前时间开始计算的有效期，类型为 time.Duration，例如 time.Hour 或 30*time.Minute。
//     ttl > 0 时设置 exp；ttl <= 0（包括 0 和任意负数）时不设置 exp，即没有过期时间，也不会因此报错。
//     不会自动使用默认有效期；如需 7 天，请显式传入 DefaultTokenTTL。
//
// JWT 的时间戳精度为秒，过期时间的小数秒部分会被截断。
func GenerateSubToken(key []byte, method *jwt.SigningMethodHMAC, user, channel string, ttl time.Duration) (string, error) {
	if channel == "" {
		return "", errKit.Newf("centrifugoKit: channel is required")
	}
	return generateToken(key, method, user, channel, ttl)
}

func generateToken(key []byte, method *jwt.SigningMethodHMAC, user, channel string, ttl time.Duration) (string, error) {
	method = cmp.Or(method, jwt.SigningMethodHS256) // 默认值

	switch method {
	case jwt.SigningMethodHS256, jwt.SigningMethodHS384, jwt.SigningMethodHS512:
	default:
		return "", errKit.Newf("centrifugoKit: signing method must be jwt.SigningMethodHS256, jwt.SigningMethodHS384 or jwt.SigningMethodHS512")
	}
	if len(key) == 0 {
		return "", errKit.Newf("centrifugoKit: signing key must not be empty")
	}

	now := time.Now()
	claims := jwt.MapClaims{"iat": now.Unix()}
	if user != "" {
		claims["sub"] = user
	}
	if channel != "" {
		claims["channel"] = channel
	}
	if ttl > 0 {
		claims["exp"] = now.Add(ttl).Unix()
	}

	token, err := jwt.NewWithClaims(method, claims).SignedString(key)
	if err != nil {
		return "", errKit.Wrapf(err, "centrifugoKit: failed to sign token")
	}
	return token, nil
}
