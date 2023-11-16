package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	t := table.Table{}
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t.AppendRows([]table.Row{{"1", "Arya", "Stark", "3000"}})
	t.AppendRows([]table.Row{{"2", "Jon", "Snow", "2000"}})
	fmt.Println(t.Render())

}
