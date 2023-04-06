package fileMode

import "os"

/*
参考：https://blog.csdn.net/qq_39131177/article/details/85060694

！！！：0开头，即八进制.

r: 读权限(4)
w: 写权限(2)
x: 执行权限(1)
-: 无权限(0)
*/
const (
	// FileModeA -rw-------: 只有拥有者有读写权限
	FileModeA = os.FileMode(0600)

	// FileModeB -rw-r--r--: 只有拥有者有读写权限；而属组用户和其他用户只有读权限
	FileModeB = os.FileMode(0644)

	// FileModeC -rwx------: 只有拥有者有读、写、执行权限
	FileModeC = os.FileMode(0700)

	// FileModeD -rwxr-xr-x: 拥有者有读、写、执行权限；而属组用户和其他用户只有读、执行权限
	FileModeD = os.FileMode(0755)

	// FileModeE -rwx--x--x: 拥有者有读、写、执行权限；而属组用户和其他用户只有执行权限
	FileModeE = os.FileMode(0711)

	// FileModeAllReadWrite -rw-rw-rw-: 所有用户 有 读、写权限
	FileModeAllReadWrite = os.FileMode(0666)

	// FileModeAll -rwxrwxrwx: 所有用户 有 读、写、执行权限
	FileModeAll = os.FileMode(0777)
)
