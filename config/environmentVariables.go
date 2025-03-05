package config

import (
	"log"
	"os"
	"strconv"
)

var (
	timeAdditionValue       int
	timeSubtractionValue    int
	timeMultiplicationValue int
	timeDivisionValue       int
	computingPowerValue     int
)

func LoadEnvironmentVariables() {
	loadTimeAddition()
	loadTimeSubtraction()
	loadTimeMultiplication()
	loadTimeDivision()
	loadComputingPower()
	log.Println("Environment Variables loaded successfully")
	log.Printf("%v; %v; %v; %v; %v; ", timeAdditionValue, timeSubtractionValue, timeMultiplicationValue, timeDivisionValue, computingPowerValue)
}

func loadTimeAddition() {
	envVal, exists := os.LookupEnv(time_addition_name)
	if !exists {
		timeAdditionValue = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeAddition() -> enviroment value=%v, is not a string", timeAdditionValue)
			timeAdditionValue = stockTimeExecution
		}
		timeAdditionValue = numVal
	}
}
func loadTimeSubtraction() {
	envVal, exists := os.LookupEnv(time_subtraction_name)
	if !exists {
		timeSubtractionValue = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeSubtraction() -> enviroment value=%v, is not a string", timeSubtractionValue)
			timeSubtractionValue = stockTimeExecution
		} else {
			timeSubtractionValue = numVal
		}
	}
}
func loadTimeMultiplication() {
	envVal, exists := os.LookupEnv(time_multiplication_name)
	if !exists {
		timeMultiplicationValue = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeMultiplication() -> enviroment value=%v, is not a string", timeMultiplicationValue)
			timeMultiplicationValue = stockTimeExecution
		} else {
			timeMultiplicationValue = numVal
		}
	}
}
func loadTimeDivision() {
	envVal, exists := os.LookupEnv(time_division_name)
	if !exists {
		timeDivisionValue = stockTimeExecution
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadTimeDivision() -> enviroment value=%v, is not a string", timeDivisionValue)
			timeDivisionValue = stockTimeExecution
		} else {
			timeDivisionValue = numVal
		}
	}
}
func loadComputingPower() {
	envVal, exists := os.LookupEnv(computing_power_name)
	if !exists {
		computingPowerValue = stockComputingPower
	} else {
		numVal, err := strconv.Atoi(envVal)
		if err != nil {
			log.Printf("ERROR: loadComputingPower() -> enviroment value=%v, is not a string", computingPowerValue)
			computingPowerValue = stockComputingPower
		} else {
			computingPowerValue = numVal
		}
	}
}

func GetTimeAddition() int {
	return timeAdditionValue
}
func GetTimeSubtraction() int {
	return timeSubtractionValue
}
func GetTimeMultiplication() int {
	return timeMultiplicationValue
}
func GetTimeDivision() int {
	return timeDivisionValue
}
func GetComputingPower() int {
	return computingPowerValue
}
