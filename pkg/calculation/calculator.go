package calculation

import (
	"strconv"
)

// Calc - Вычисляет математическое выражение, переданное в виде строки.
// Функция использует алгоритм RPN
func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0, ErrEmptyExpression
	}

	tokens := tokenize(expression)
	postfixExpression, err := infixToPostfix(tokens)
	if err != nil {
		return 0, err
	}
	result, err := evaluatePostfixExpression(postfixExpression)
	return result, err
}

// evaluatePostfixExpression - Вычисление постфиксного выражения
func evaluatePostfixExpression(postfixTokens []string) (float64, error) {
	var stack []float64

	for _, token := range postfixTokens {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, ErrInvalidExpression
			}
			stack = append(stack, num)
		} else if isOperator(token) {
			if len(stack) < 2 {
				return 0, ErrInvalidExpression
			}
			{
				a, b := stack[len(stack)-2], stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				result, err := calculateSimple(a, b, token)
				if err != nil {
					return 0, err
				}
				stack = append(stack, result)
			}
		} else {
			return 0, ErrInvalidExpression
		}
	}
	if len(stack) != 1 {
		return 0, ErrInvalidExpression
	}
	return stack[0], nil
}
