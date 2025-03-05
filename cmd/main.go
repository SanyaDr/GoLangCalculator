package main

import (
	"SecondSprintExam/config"
	"SecondSprintExam/internal/app"
	"log"
	"os"
)

// Run функции для запуска http сервера
// TODO изучить godoc
// TODO добавить api тесты
// TODO переписать на нормальный logger :(
func setLogger() {
	log.SetOutput(os.Stdout)
}

// initVars - инициализация необходимых переменных
func initVars() {
	config.LoadEnvironmentVariables()
}

func main() {
	setLogger()
	initVars()
	app.RunServer()
}
