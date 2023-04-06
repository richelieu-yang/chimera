package sliceKit

type (
	// RemoveJudge 判断是否将元素从slice中移除
	/*
		@param 	ele			slice中的1个元素
		@return	remove		是否将该元素移除？
		@return	interrupt	是否中断遍历？
	*/
	RemoveJudge[T comparable] func(ele T) (remove bool, interrupt bool)
)
