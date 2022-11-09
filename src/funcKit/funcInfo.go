package funcKit

// GetFuncInfo
/*
@param extraSkips 额外跳过的步骤，第1个值（有的话）必须: >= 0
@return extraSkip == 0的情况下，返回: 调用此函数的函数的信息（"$包名.$方法名: "）

e.g.
"main.main: "
"main.Print: "
"main.test: "
*/
func GetFuncInfo(extraSkips ...int) string {
	var extraSkip int
	if extraSkips != nil {
		extraSkip = extraSkips[0]
	} else {
		extraSkip = 0
	}

	return GetCallerNameWithSkip(2+extraSkip) + ": "
}

// AddFuncInfoToString 在传参format前面加上: "$包名.$方法名: "
/*
@param extraSkips 额外跳过的步骤，第1个值（有的话）必须: >= 0
*/
func AddFuncInfoToString(str string, extraSkips ...int) string {
	var extraSkip int
	if extraSkips != nil {
		extraSkip = extraSkips[0]
	} else {
		extraSkip = 0
	}

	// 此处+1的原因: 多了一层调用
	return GetFuncInfo(extraSkip+1) + str
}
