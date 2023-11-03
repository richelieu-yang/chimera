package types

import "net/http"

type Listeners []Listener

func (listeners Listeners) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	for _, listener := range listeners {
		listener.OnFailure(w, r, failureInfo)
	}
}

func (listeners Listeners) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	for _, listener := range listeners {
		listener.OnHandshake(w, r, channel)
	}
}

func (listeners Listeners) OnMessage(channel Channel, messageType int, data []byte) {
	for _, listener := range listeners {
		listener.OnMessage(channel, messageType, data)
	}
}

func (listeners Listeners) OnClose(channel Channel, closeInfo string) {
	for _, listener := range listeners {
		listener.OnClose(channel, closeInfo)
	}
}
