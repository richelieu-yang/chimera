package fileMode

import "os"

/*
参考：https://blog.csdn.net/qq_39131177/article/details/85060694

！！！：0开头，即八进制.

-: 无权限
r: 读权限
w: 写权限
x: 执行权限
*/
const (
	// ModeA
	// -rw-------: 只有拥有者有读写权限
	ModeA = os.FileMode(0600)

	// ModeB
	// -rw-r--r--: 只有拥有者有读写权限；而属组用户和其他用户只有读权限
	ModeB = os.FileMode(0644)

	// ModeC
	// -rwx------: 只有拥有者有读、写、执行权限
	ModeC = os.FileMode(0700)

	// ModeD
	// -rwxr-xr-x: 拥有者有读、写、执行权限；而属组用户和其他用户只有读、执行权限
	ModeD = os.FileMode(0755)

	// ModeE
	// -rwx--x--x: 拥有者有读、写、执行权限；而属组用户和其他用户只有执行权限
	ModeE = os.FileMode(0711)

	// ModeAllReadWrite
	// -rw-rw-rw-: 所有用户 有 读、写权限
	ModeAllReadWrite = os.FileMode(0666)

	// ModeAll
	// -rwxrwxrwx: 所有用户 有 读、写、执行权限
	ModeAll = os.FileMode(0777)
)
