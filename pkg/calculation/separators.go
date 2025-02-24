package calculator

import (
	"strconv"
	"strings"
)

// Возвращает слайс чисел из строки
func separateNumbers(inp string) ([]float64, error) {
	numbersStr := strings.FieldsFunc(inp, separatorOperator)
	numbers := make([]float64, len(numbersStr))
	for i, s := range numbersStr {
		var err error
		numbers[i], err = strconv.ParseFloat(s, 64)
		if err != nil {
			return numbers, ErrInvalidExpression
		}
	}
	return numbers, nil
}

// Возвращает слайс операторов из строки
func separateOperators(inp string) []rune {
	var operators []rune
	for _, v := range inp {
		if strings.ContainsRune(operatorChars, v) {
			operators = append(operators, v)
		}
	}
	return operators
}
