package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

// Initialize 初始化
/*
@param upgrader	可以为nil
@param listener	可以为nil
@param logger	可以为nil，默认直接输出到控制台
*/
func Initialize(upgrader *websocket.Upgrader, listener Listener, logger *logrus.Logger) error {
	setUpgrader(upgrader)
	setListener(listener)

	return nil
}
