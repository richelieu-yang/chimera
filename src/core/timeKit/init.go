package timeKit

//// CSTLocation CST: 中国标准时间(China Standard Time)
///*
//！！！：可能为nil.
//e.g.
//Windows系统，如果安装了Golang环境，能获取到；否则获取不到（error: The system cannot find the path specified.）.
//*/
//var CSTLocation *time.Location
//
//func init() {
//	// LoadLocation的输入参数的取值，除了该函数的源代码中可看到的”UTC”、”Local”，其余的值其实是遵照“IANA Time Zone”的规则，可以解压$GOROOT/lib/time/zoneinfo.zip 这个文件打开查看。
//	// 在Asia这个目录，我看到了Chongqing，Hong_Kong，但没Beijing。在国外获取中国北京时间，要用”PRC”，当然”Asia/Chongqing”也是个方法
//	// 参考：https://blog.csdn.net/qq_26981997/article/details/53454606
//	names := []string{"Asia/Shanghai", "Asia/Chongqing"}
//	for _, name := range names {
//		if loc, err := LoadLocation(name); err != nil {
//			// e.g. 获取不到
//			// error: The system cannot find the path specified.
//			logrus.Warnf("[SCALES, TIME] fail to load time location(name: %s), error: %v", name, err)
//		} else {
//			CSTLocation = loc
//			break
//		}
//	}
//}
