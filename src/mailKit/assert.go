package mailKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
	"github.com/richelieu-yang/chimera/v2/src/validateKit"
)

func AssertEmail(email string) error {
	if err := validateKit.Email(email); err != nil {
		return errorKit.NewSkip(1, "[%s] email(%s) is invalid with error(%s)", funcKit.GetFuncName(1), email, err.Error())
	}
	return nil
}
