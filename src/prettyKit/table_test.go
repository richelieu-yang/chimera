package prettyKit

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"testing"
)

/*
output:
+---+------------+-----------+--------+
| # | FIRST NAME | LAST NAME | SALARY |
+---+------------+-----------+--------+
| 1 | Arya       | Stark     | 3000   |
| 2 | Jon        | Snow      | 2000   |
+---+------------+-----------+--------+
*/
func TestNewTableWriter(t *testing.T) {
	writer := NewTableWriter()
	writer.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	writer.AppendRows([]table.Row{{"1", "Arya", "Stark", "3000"}})
	writer.AppendRows([]table.Row{{"2", "Jon", "Snow", "2000"}})
	fmt.Println(writer.Render())
}

func TestCreateTable(t *testing.T) {
	table := CreateTable()
	table.AddHeaders("User", "Age")
	table.AddRow("San Zhang", 18)
	table.AddRow("Si Li", 30)

	fmt.Println("默认格式（普通的文本表格）:")
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
