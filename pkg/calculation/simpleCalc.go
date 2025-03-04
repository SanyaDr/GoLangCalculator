package calculator

import (
	outer "SecondSprintExam/config"
	"strings"
	"time"
)

var operatorChars = "+-*/" // Поддерживаемые операции
var separatorOperator = func(r rune) bool {
	return strings.ContainsRune(operatorChars, r)
}

// Функция возвращает результат заданной операции
// a, b float64 - переменные для вычисления
// operator rune - знак операции (+, -, *, /)
func calculate(a, b float64, operation rune) (float64, error) {
	switch operation {
	case '+':
		<-time.After(time.Duration(outer.GetTimeAddition()) * time.Millisecond)
		return a + b, nil
	case '-':
		<-time.After(time.Duration(outer.GetTimeSubtraction()) * time.Millisecond)
		return a - b, nil
	case '*':
		<-time.After(time.Duration(outer.GetTimeMultiplication()) * time.Millisecond)
		return a * b, nil
	case '/':
		<-time.After(time.Duration(outer.GetTimeDivision()) * time.Millisecond)
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil

	default:
		return 0, ErrUnsupportedOperation
	}
}

// Выполняет простые вычисления без скобок, например: 2+3*4/5
func calculateSimple(inp string) (float64, error) {
	numbers, err := separateNumbers(inp) // Слайс чисел
	operators := separateOperators(inp)  // Слайс операторов
	if err != nil {
		return 0, err
	}
	// Если операторов больше чем чисел, то ошибка
	if len(operators) >= len(numbers) {
		return 0, ErrInvalidExpression
	}

	// Если есть приоритетные операторы (умножения, деления)
	if strings.ContainsAny(inp, "*/") {
		for i := 0; i < len(operators); i++ {
			if operators[i] == '*' || operators[i] == '/' {
				numbers[i], err = calculate(numbers[i], numbers[i+1], operators[i])
				if err != nil {
					return 0, err
				}
				numbers = append(numbers[:i+1], numbers[i+2:]...)
				operators = append(operators[:i], operators[i+1:]...)
				i--
			}
		}
	}
	//если приоритетных операций больше нет
	for len(numbers) != 1 {
		numbers[0], err = calculate(numbers[0], numbers[1], operators[0])
		if err != nil {
			return 0, err
		}
		numbers = append(numbers[:1], numbers[2:]...)
		operators = operators[1:]
	}
	return numbers[0], nil
}
