package calculator

import (
	"fmt"
	"strings"
)

func simplifyParentheses(expression string) (string, error) {
	parOpenCount, parCloseCount := 0, 0
	parFirstOpen, parLastClose := 0, 0

	for i := 0; i < len(expression); i++ {
		cur := expression[i]
		if cur == '(' {
			parOpenCount++
			if parOpenCount == 1 {
				parFirstOpen = i
			}
		}
		if cur == ')' {
			parCloseCount++
			if parCloseCount == parOpenCount {
				parLastClose = i

				simpleExp, err := simplifyParentheses(expression[parFirstOpen+1 : parLastClose])
				if err != nil {
					return "", ErrInvalidExpression
				}
				num, err := calculateSimple(simpleExp)
				if err != nil {
					return "", err
				}
				i -= parLastClose - parFirstOpen
				expression = expression[:parFirstOpen] + fmt.Sprintf("%v", num) + expression[parLastClose+1:]

				parOpenCount, parCloseCount = 0, 0
			}
		}
	}
	if parOpenCount != parCloseCount {
		return "", ErrMismatchedParentheses
	}

	return expression, nil
}

// Убрать пробелы
func simplifySpaces(inp string) string {
	result := strings.ReplaceAll(inp, " ", "")
	return result
}
