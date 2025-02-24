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

func main() {
	setLogger()
	app.RunServer()
}
