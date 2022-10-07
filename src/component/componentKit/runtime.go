package componentKit

import (
	"github.com/richelieu42/go-scales/src/core/runtimeKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/dataSizeKit"
	"github.com/sirupsen/logrus"
)

type (
	RuntimeConfig struct {
		// 软内存限制（如果remove space后为""或"-1"，将不进行设置）
		SoftMemoryLimit string
	}
)

// initializeRuntimeComponent 初始化 runtime 组件
func initializeRuntimeComponent() error {
	config, err := GetRuntimeConfig()
	if err != nil {
		return err
	}
	//if config == nil {
	//	return errorKit.Simple("config == nil")
	//}

	/* SoftMemoryLimit */
	if config == nil {
		logrus.Info("[COMPONENT, RUNTIME] softMemoryLimit won't be set.")
	} else {
		limit := config.SoftMemoryLimit
		limit = strKit.RemoveSpace(limit)
		switch limit {
		case "":
			fallthrough
		case "-1":
			logrus.Info("[COMPONENT, RUNTIME] softMemoryLimit won't be set.")
		default:
			size, err := dataSizeKit.ParseStringToDataSize(limit)
			if err != nil {
				return err
			}
			byteValue := int64(size.GetByteValue())
			if _, err := runtimeKit.SetSoftMemoryLimit(byteValue); err != nil {
				return err
			}
			logrus.Infof("[COMPONENT, RUNTIME] softMemoryLimit: [%s].", size.ToString())
		}
	}

	logrus.Info("[COMPONENT, RUNTIME] Initialize successfully.")
	return nil
}
