package main

import (
	"SecondSprintExam/config"
	"SecondSprintExam/internal/app"
	"log"
	"os"
)

// Run фукнции для запуска http сервера
// TODO изучить godoc
// TODO добавить api тесты
// TODO переписать на нормальный logger :(
// TODO Добавь в readme возможность изменить порт запуска в конфиге
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
