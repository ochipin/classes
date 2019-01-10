package main

var globalValue int

func GlobalFunction() {
}

type App struct {
	Name string
}

func (app *App) Main() {
}

type base struct {
	age int
}

func (b base) Main() {
}
