package main

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	t := table.Table{}
	t.AppendHeader(table.Row{"#", "First Name", "Last Name", "Salary"})
	t.AppendRows([]table.Row{{"1", "Arya", "Stark", "3000"}})
	fmt.Println(t.Render())

}
