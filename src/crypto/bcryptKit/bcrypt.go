package bcryptKit

import "golang.org/x/crypto/bcrypt"

// HashPassword 加密用户密码（生成一个带盐（salt）的哈希值）.
/*
PS: 传参相同，多次执行的结果不同.
*/
func HashPassword(password []byte) ([]byte, error) {
	// 迭代次数，可以根据需要调整
	cost := bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	return hashedPassword, err
}

// ComparePasswords 验证密码是否正确（对比 已哈希过的密码 与 原始明文密码 是否匹配。）.
/*
@param hashedPassword HashPassword() 的返回值
@param plainPassword 密码明文
@return err == nil: 已哈希过的密码 与 原始明文密码 匹配
*/
func ComparePasswords(hashedPassword, plainPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)
}
