package logrusKit

/*
初始化 logrus 相关的配置.

PS:
(1) 可以在 main() 所在的.go文件中，先通过 _ 导入此包，进行初始化配置；
(2) 后续（比如读取完配置文件）可以重新调用 MustSetUp().
*/
func init() {
	MustSetUp(nil)
}
