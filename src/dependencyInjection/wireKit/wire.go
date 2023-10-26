package wireKit

import "github.com/google/wire"

var (
	// Bind 建立接口类型和具体的实现类型之间的绑定关系，这样 Wire 工具就可以根据这个绑定关系进行类型匹配并生成代码.
	/*
		wire.Bind 的第一个参数是指向所需接口类型值的指针，第二个实参是指向实现该接口的类型值的指针。
	*/
	Bind func(iface, to interface{}) wire.Binding = wire.Bind

	// Struct 根据现有的类型进行构造结构体.
	/*
		使用 wire.Struct 函数需要传递两个参数，第一个参数是结构体类型的指针值，另一个参数是一个可变参数，表示需要注入的结构体字段的名称集。
	*/
	Struct func(structType interface{}, fieldNames ...string) wire.StructProvider = wire.Struct

	// Value
	Value = wire.Value

	// InterfaceValue
	InterfaceValue = wire.InterfaceValue

	// FieldsOf
	FieldsOf = wire.FieldsOf
)
