package tableKit

import (
	"fmt"
	"testing"
)

func TestCreateTable(t *testing.T) {
	table := CreateTable()
	table.AddHeaders("User", "Age")
	table.AddRow("San Zhang", 18)
	table.AddRow("Si Li", 30)

	fmt.Println("默认格式:")
	fmt.Println(table.Render())

	fmt.Println("HTML格式:")
	table.SetModeHTML()
	fmt.Println(table.Render())

	fmt.Println("Markdown格式:")
	table.SetModeMarkdown()
	fmt.Println(table.Render())
}

// 可以只调用 Table.AddRow()
func TestCreateTable1(t *testing.T) {
	table := CreateTable()
	//table.AddHeaders("User", "Age")
	table.AddRow("San Zhang", 18)
	table.AddRow("Si Li", 30)

	fmt.Println(table.Render())
}
