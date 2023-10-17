package main

import "github.com/richelieu-yang/chimera/v2/src/validateKit"

type User struct {
	Email    string `validate:"email"`
	Password string `validate:"required"`
	Token    string `validate:"-"`
}

func (u *User) Validate() error {
	v := validateKit.New()

	// 如果Token为空，则验证 Email 和 Password
	if u.Token == "" {
		return v.Struct(u)
	}
	// 否则只验证 Token
	return v.Var(u.Token, "required")
}

func main() {
	u := &User{
		Email:    "",
		Password: "",
		Token:    "111",
	}

	v := validateKit.New()
	if err := v.Struct(u); err != nil {
		panic(err)
	}
}
