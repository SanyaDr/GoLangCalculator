package SecondSprintExam

import (
	"log"
	"os"
	"strconv"
)

// Названия переменных окружения системы
const (
	stockTimeExecution  = 10 // Время выполнение операции по умолчанию (если переменные среды не найдены)
	stockComputingPower = 3  // Кол-во горутин по умолчанию

	time_addition_name       = "TIME_ADDITION_MS"       // Время выполнения операции сложения в миллисекундах
	time_subtraction_name    = "TIME_SUBTRACTION_MS "   // Время выполнения операции вычитания в миллисекундах
	time_multiplication_name = "TIME_MULTIPLICATION_MS" // Время выполнения операции умножения в миллисекундах
	time_division_name       = "TIME_DIVISION_MS"       // Время выполнения операции деления в миллисекундах
	computing_power_name     = "COMPUTING_POWER"        // Количество горутин регулируется переменной среды
)

var (
	time_addition_value       int
	time_subtraction_value    int
	time_multiplication_value int
	time_division_value       int
	computing_power_value     int
)

func LoadEnvironmentVariables() {
	loadTimeAddition()
	loadTimeSubtraction()
	loadTimeMultiplication()
	loadTimeDivision()
	loadComputingPower()
	log.Println("Environment Variables loaded successfully")
}

func loadTimeAddition() {
	envVal, exists := os.LookupEnv(time_addition_name)
	if !exists {
		time_addition_value = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeAddition() -> enviroment value=%v, is not a string", time_addition_value)
			time_addition_value = stockTimeExecution
		}
		time_addition_value = numVal
	}
}
func loadTimeSubtraction() {
	envVal, exists := os.LookupEnv(time_subtraction_name)
	if !exists {
		time_subtraction_value = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeSubtraction() -> enviroment value=%v, is not a string", time_subtraction_value)
			time_subtraction_value = stockTimeExecution
		} else {
			time_addition_value = numVal
		}
	}
}
func loadTimeMultiplication() {
	envVal, exists := os.LookupEnv(time_multiplication_name)
	if !exists {
		time_multiplication_value = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeMultiplication() -> enviroment value=%v, is not a string", time_multiplication_value)
			time_multiplication_value = stockTimeExecution
		} else {
			time_multiplication_value = numVal
		}
	}
}
func loadTimeDivision() {
	envVal, exists := os.LookupEnv(time_division_name)
	if !exists {
		time_division_value = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeDivision() -> enviroment value=%v, is not a string", time_division_value)
			time_division_value = stockTimeExecution
		} else {
			time_division_value = numVal
		}
	}
}
func loadComputingPower() {
	envVal, exists := os.LookupEnv(computing_power_name)
	if !exists {
		computing_power_value = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadComputingPower() -> enviroment value=%v, is not a string", computing_power_value)
			computing_power_value = stockTimeExecution
		} else {
			computing_power_value = numVal
		}
	}
}

func GetTimeAddition() int {
	return time_addition_value
}
func GetTimeSubtraction() int {
	return time_subtraction_value
}
func GetTimeMultiplication() int {
	return time_multiplication_value
}
func GetTimeDivision() int {
	return time_division_value
}
func GetComputingPower() int {
	return computing_power_value
}
