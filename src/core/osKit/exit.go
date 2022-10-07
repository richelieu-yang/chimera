package osKit

import (
	"os"
)

// ExitWithCode 退出程序
/*
PS: 无论是在main程还是子程中，只要调用os.Exit()，程序就会终止.

@param code 0：正常退出；非0：非正常退出（一般用1）
*/
func ExitWithCode(code int) {
	os.Exit(code)
}
