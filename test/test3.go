package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `
3个码：2BTE7KLR4PEQ兑换码4ATWKLRTPZWU兑换码QBBSZ57HHS24
2个码：DBSGA96YL3YQ MAT7Y6NJQP2Y
5个兑换码烦请分别兑换
可+微上新第一时间通知：sanqiangame
游戏内>左上角菜单>右上角...>兑换码
如需本店更多星铁商品请点击：https://s.tb.cn/c.0vdwpI
`

	{
		re := regexp.MustCompile("[a-zA-Z0-9]{12,14}")
		s := re.FindAllString(str, -1)
		fmt.Println(len(s)) // 5
		fmt.Println(s)      // [2BTE7KLR4PEQ 4ATWKLRTPZWU QBBSZ57HHS24 DBSGA96YL3YQ MAT7Y6NJQP2Y]
	}
}
