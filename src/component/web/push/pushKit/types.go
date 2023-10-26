package pushKit

import "net/http"

type Listener interface {
	OnFailure(w http.ResponseWriter, r *http.Request, error string)
}

type Channel interface {
	// Push 推送消息给客户端
	Push() error
}
