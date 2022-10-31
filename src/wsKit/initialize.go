package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
)

var logDir = "."

// Initialize 初始化
/*
@param upgrader	可以为nil
@param listener	可以为nil
@param logDir1	输出日志的目录
*/
func Initialize(upgrader *websocket.Upgrader, listener Listener, logDir1 string) error {
	setUpgrader(upgrader)
	setListener(listener)

	/* logDir */
	if strKit.IsNotEmpty(logDir1) {
		logDir = logDir1
	}
	if err := fileKit.MkDirs(logDir); err != nil {
		return err
	}

	return nil
}
