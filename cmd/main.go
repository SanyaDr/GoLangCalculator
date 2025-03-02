package main

import (
	"SecondSprintExam/internal/app"
	"log"
	"os"
)

// Run фукнции для запуска http сервера
// TODO изучить godoc
// TODO добавить api тесты
func setLogger() {
	log.SetOutput(os.Stdout)
}

// initVars - инициализация необходимых переменных
func initVars() {
	//TODO удали, делай сразу через = make...
	//app.NewExpressionsStorage()
}

func main() {
	setLogger()
	initVars()
	app.RunServer()
}
