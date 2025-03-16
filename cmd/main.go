package main

import (
	"GoLangCalculator/config"
	calculator "GoLangCalculator/pkg/calculation"
	"fmt"
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
	num, err := calculator.Calc("(2)+(3)")
	fmt.Printf("num: %v; err: %v\n", num, err)

	//setLogger()
	//initVars()
	//app.RunServer()
}
