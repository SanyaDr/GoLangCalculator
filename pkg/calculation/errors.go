package calculator

import "errors"

var (
	ErrDivisionByZero        = errors.New("division by zero")                   // Деление на 0
	ErrMismatchedParentheses = errors.New("mismatched parentheses")             // Пропущена круглая скобка в выражении
	ErrInvalidExpression     = errors.New("invalid expression")                 // Недопустимое выражение. Пропущен знак или цифра
	ErrEmptyExpression       = errors.New("empty expression")                   // Пустое выражение
	ErrUnsupportedOperation  = errors.New("this operation sign is unsupported") // Данная операция не поддерживается
)
