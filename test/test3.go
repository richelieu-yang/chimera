package main

//import (
//	"github.com/go-playground/validator/v10"
//)
//
//type User struct {
//	//Status string `validate:"eq=active"`
//	Status string `validate:"eq=active"`
//	Age    int    `validate:"required_if=Status active"`
//}
//
//// 自定义验证函数
//func checkCondition(fl validator.FieldLevel) bool {
//	if fl.Field().String() == "active" {
//		return fl.Parent().FieldByName("Age").IsValid()
//	}
//	return true
//}
//
//func main() {
//	validate := validator.New()
//
//	// 将自定义函数checkCondition与 struct tag 关联起来
//	err := validate.RegisterValidation("required_if", checkCondition)
//	if err != nil {
//		// 处理错误
//	}
//
//	// 现在你可以在你的struct tag中使用"required_if"作为验证规则了
//}
