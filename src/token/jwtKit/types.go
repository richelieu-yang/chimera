package jwtKit

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

type JWT struct {
	Key []byte
	// shift 偏移量
	shift int
}

func NewJWT(key []byte, shiftArgs ...int) *JWT {
	return &JWT{
		Key:   key,
		shift: sliceKit.GetFirstItemWithDefault(0, shiftArgs...),
	}
}

func (j *JWT) check() error {
	if err := interfaceKit.AssertNotNil(j, "j"); err != nil {
		return err
	}
	return sliceKit.AssertNotEmpty(j.Key, "Key")
}

// New
/*
@param Key		密钥
@param method 	e.g. jwt.SigningMethodHS256 || jwt.SigningMethodHS384 || jwt.SigningMethodHS512
*/
func (j *JWT) New(method jwt.SigningMethod, claims jwt.MapClaims, options ...jwt.TokenOption) (string, error) {
	if err := j.check(); err != nil {
		return "", err
	}
	if err := interfaceKit.AssertNotNil(method, "method"); err != nil {
		return "", err
	}
	if err := mapKit.AssertNotEmpty(claims, "claims"); err != nil {
		return "", err
	}

	/*
		Create a new token object, specifying signing method and the claims
		you would like it to contain.
	*/
	token := jwt.NewWithClaims(method, claims, options...)

	/*
		Sign and get the complete encoded token as a string using the secret
	*/
	return token.SignedString(j.Key)
}

// Parse
/*
@param keyFunc e.g.
	func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return Key, nil
	}
*/
func (j *JWT) Parse(tokenString string, keyFunc jwt.Keyfunc, options ...jwt.ParserOption) (jwt.MapClaims, error) {
	if err := j.check(); err != nil {
		return nil, err
	}
	if err := strKit.AssertNotEmpty(tokenString, "tokenString"); err != nil {
		return nil, err
	}
	if len(strKit.Split(tokenString, ".")) != 3 {
		return nil, errorKit.New("tokenString(%s) is invalid", tokenString)
	}
	if err := interfaceKit.AssertNotNil(keyFunc, "keyFunc"); err != nil {
		return nil, err
	}

	// Parse takes the token string and a function for looking up the Key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which Key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, keyFunc, options...)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errorKit.New("token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errorKit.New("type(%T) of claims is invalid", token.Claims)
	}
	return claims, nil
}
