package consoleKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/richelieu-yang/chimera/v2/src/cronKit"
)

func RotateConsoleOutput(output, backDir, spec string) error {
	c, _, err := cronKit.NewCronWithTask(spec, func() {
		if err := fileKit.AssertExistAndIsFile(output); err != nil {

			return
		}

	})
	if err != nil {
		return err
	}
	return nil
}
