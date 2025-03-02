package main

import (
	envi "SecondSprintExam"
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
	envi.LoadEnvironmentVariables()
}

func main() {
	setLogger()
	initVars()
	app.RunServer()
}
