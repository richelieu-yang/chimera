package viperKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

// MarshalToFile
/*
@param obj 如果为nil，将返回error
*/
func MarshalToFile(obj interface{}, target string) error {
	if err := interfaceKit.AssertNotNil(obj, "obj"); err != nil {
		return err
	}
	if err := fileKit.AssertNotExistOrIsFile(target); err != nil {
		return err
	}

	data, err := jsonKit.Marshal(obj)
	if err != nil {
		return err
	}
	v, err := Read(data, "json", nil)
	if err != nil {
		return err
	}
	return v.WriteConfigAs(target)
}
