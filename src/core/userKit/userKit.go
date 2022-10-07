package userKit

func GetName() string {
	return u.Name
}

func GetUserName() string {
	return u.Username
}

// GetHomeDir 获取当前用户的home dir
/*
也可以通过 第三方库go-homedir（https://github.com/mitchellh/go-homedir） 来获取，但感觉没有必要.

@return e.g."/Users/richelieu"
*/
func GetHomeDir() string {
	return homeDir
}
