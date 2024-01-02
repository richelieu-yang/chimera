package httpKit

import "net/http"

func GetReferer(r *http.Request) string {
	return r.Referer()
}
