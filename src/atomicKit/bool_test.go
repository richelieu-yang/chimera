package atomicKit

import (
	"fmt"
	"testing"
)

func TestNewBool(t *testing.T) {
	b := NewBool()
	/* 默认false */
	fmt.Println("initial value:", b.Val()) // initial value: false

	/* 修改值成功，但这么干无意义（目前是false，修改为false） */
	fmt.Println(b.Cas(false, false)) // true
	fmt.Println("value:", b.Val())   // value: false

	/* 修改值失败 */
	fmt.Println(b.Cas(true, false)) // false
	fmt.Println("value:", b.Val())  // value: false

	/* 修改值成功 */
	fmt.Println(b.Cas(false, true)) // true
	fmt.Println("value:", b.Val())  // value: true
}
