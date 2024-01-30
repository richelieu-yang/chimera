package httpKit

var (
	RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP"}
	//RemoteIPHeaders = []string{"X-Forwarded-For", "X-Real-IP", "Client-IP"}
)
