package main

import "fmt"

type Database interface {
	Query() string
}

type MySQL struct{}

func (m MySQL) Query() string {
	return "MySQL Query"
}

type App struct {
	db Database
}

func NewApp(db Database) *App {
	return &App{db: db}
}

func (a App) Run() {
	result := a.db.Query()
	fmt.Println(result)
}

func main() {
	mySQL := MySQL{}
	app := NewApp(mySQL)
	app.Run()
}
