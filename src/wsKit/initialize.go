package wsKit

import "github.com/gorilla/websocket"

// Initialize 初始化
/*
@param upgrader	可以为nil
@param listener	可以为nil
*/
func Initialize(upgrader *websocket.Upgrader, listener Listener) error {
	setUpgrader(upgrader)
	setListener(listener)

	return nil
}
