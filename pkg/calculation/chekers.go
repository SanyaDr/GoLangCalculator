package calculation

import "strconv"

// func isNumber, isOperator,getPrecedence итд
//
//// IsParseFloat64 - Если переданная строка число, возвращает число и true, иначе 0 и false
//func isParseFloat64(inputNumber string) (float64, bool) {
//	num, err := strconv.ParseFloat(inputNumber, 64)
//	if err != nil {
//		return 0, false
//	}
//	return num, true
//}

// IsNumber - проверяет, является ли строка числом
func isNumber(inp string) bool {
	_, err := strconv.ParseFloat(inp, 64)
	return err == nil
}

// IsOperator - возвращает, является ли полученный токен поддерживаемым оператором
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "*" || token == "/"
}

// getPrecedence - возвращает приоритет оператора
func getPrecedence(operator string) int {
	switch operator {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	default:
		return 0
	}
}
