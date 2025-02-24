package calculator

// Calc - Входная точка в калькулятор.
// Принимает выражение в виде строки.
//
// Возвращает результат и ошибку в формате: float64 (равен нулю в случае ошибки), error (равен nil если ошибок нет)
func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0, ErrEmptyExpression
	}
	expression = simplifySpaces(expression)

	simple, err := simplifyParentheses(expression)
	if err != nil {
		return 0, err
	}
	res, err := calculateSimple(simple)
	if err != nil {
		return 0, err
	}

	return res, nil
}
