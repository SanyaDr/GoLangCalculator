package main

import (
	"log"
	"os"
)

// Названия переменных окружения системы
const (
	stockTimeExecution = "1" // Время выполнение операции по умолчанию (если переменные среды не найдены)

	time_addition_name       = "TIME_ADDITION_MS"
	time_subtraction_name    = "TIME_SUBTRACTION_MS "
	time_multiplication_name = "TIME_MULTIPLICATION_MS"
	time_division_name       = "TIME_DIVISION_MS"
	computing_power_name     = "COMPUTING_POWER"

	//TIME_ADDITION_MS        = 10 // Время выполнения операции сложения в миллисекундах
	//TIME_SUBTRACTION_MS     = 10 // Время выполнения операции вычитания в миллисекундах
	//TIME_MULTIPLICATIONS_MS = 10 // Время выполнения операции умножения в миллисекундах
	//TIME_DIVISIONS_MS       = 10 // Время выполнения операции деления в миллисекундах
	//COMPUTING_POWER         = 3  // Количество горутин регулируется переменной среды
)

var time_addition_value, time_subtraction_value, time_multiplication_value, time_division_value, computing_power_value int

func setEnvStockValue(name string) {
	err := os.Setenv(name, stockTimeExecution)
	if err != nil {
		log.Fatalf("FATAL: Ошибка запуска сервера! Ошибка установки значения переменной!\n %v", err)
	}
}

func checkEnvironmentVariables() {
	// TODO А что если не устанавливать переменные среды, а просто использовать переменные по умолчанию? не устанавливая новые. Проверь потом как будешь использовать это в функциях
	time_addition_value, exists := os.LookupEnv(time_addition_name)
	if !exists {
		setEnvStockValue(time_addition_name)
	}
	time_subtraction_value, exists := os.LookupEnv(time_subtraction_name)
	if !exists {
		setEnvStockValue(time_subtraction_name)
	}
	time_multiplication_value, exists := os.LookupEnv(time_multiplication_name)
	if !exists {
		setEnvStockValue(time_multiplication_name)
	}
	time_division_value, exists := os.LookupEnv(time_division_name)
	if !exists {
		setEnvStockValue(time_division_name)
	}
	computing_power_value, exists := os.LookupEnv(computing_power_name)
	if !exists {
		setEnvStockValue(computing_power_name)
	}

	log.Println("Переменные среды успешно получены!")
}
