package randomKit

//// Float64
///*
//@param places 保留的小数位
//@return 范围: [min, max)
//*/
//func Float64(min, max float64, places int32) (rst float64, err error) {
//	if min >= max {
//		err = errorKit.Newf("min(%d) is greater than or equal to max(%d)", min, max)
//		return
//	}
//
//	// rst范围: [0, max - min)
//	rst = r.Float64() * floatKit.Sub(max, min)
//	// rst范围: [min, max)
//	rst = floatKit.Add(rst, min)
//	// 保留小数位
//	rst = floatKit.Floor(rst, places)
//	return
//}
