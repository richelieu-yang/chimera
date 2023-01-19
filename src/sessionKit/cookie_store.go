package sessionKit

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// NewCookieStore
/*
PS: 以 sessions.NewCookieStore() 为基础进行修改.

@param keyPairs	[]byte("0123456789abcdef0123456789abcdef")
@param options 	可以为nil（将采用默认值）
*/
func NewCookieStore(keyPairs []byte, options *sessions.Options) *sessions.CookieStore {
	if options == nil {
		options = &sessions.Options{
			Path:   "/",
			MaxAge: 1800, /* 默认: 30min */
		}
	}

	cookie := &sessions.CookieStore{
		Codecs:  securecookie.CodecsFromPairs(keyPairs),
		Options: options,
	}
	cookie.MaxAge(cookie.Options.MaxAge)
	return cookie
}
