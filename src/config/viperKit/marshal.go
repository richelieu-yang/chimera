package viperKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

func MarshalToFile(obj interface{}, target string) error {
	if err := interfaceKit.AssertNotNil(obj, "obj"); err != nil {
		return err
	}

	data, err := jsonKit.Marshal(obj)
	if err != nil {
		return err
	}
	Unmarshal(data, "json")

	return nil
}
