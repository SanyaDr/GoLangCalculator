package calculation

import (
	"strings"
	"unicode"
)

// tokenize - разбивает строку на токены (числа и операторы).
func tokenize(expression string) []string {
	var tokens []string
	var currentToken strings.Builder

	for i, char := range expression {
		if unicode.IsSpace(char) {
			continue
		}
		if unicode.IsDigit(char) || char == '.' {
			currentToken.WriteRune(char)
		} else {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			if char == '-' {
				if i == 0 || expression[i-1] == '(' || isOperator(string(expression[i-1])) {
					currentToken.WriteRune(char)
				} else {
					tokens = append(tokens, string(char))
				}
			} else {
				tokens = append(tokens, string(char))
			}
		}
	}
	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}
	return tokens
}

// infixToPostfix - преобразование инфиксного написания в постфиксное
func infixToPostfix(tokens []string) ([]string, error) {
	var resultPostfixExpression []string
	var operators []string

	for _, token := range tokens {
		if isNumber(token) {
			resultPostfixExpression = append(resultPostfixExpression, token)
		} else if token == "(" {
			operators = append(operators, token)
		} else if token == ")" {
			for len(operators) > 0 && operators[len(operators)-1] != "(" {
				resultPostfixExpression = append(resultPostfixExpression, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			if len(operators) == 0 {
				return nil, ErrMismatchedParentheses
			}
			operators = operators[:len(operators)-1]
		} else if isOperator(token) {
			for len(operators) > 0 && getPrecedence(operators[len(operators)-1]) >= getPrecedence(token) {
				resultPostfixExpression = append(resultPostfixExpression, operators[len(operators)-1])
				operators = operators[:len(operators)-1]
			}
			operators = append(operators, token)
		} else {
			return nil, ErrInvalidExpression
		}
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == "(" {
			return nil, ErrMismatchedParentheses
		}
		resultPostfixExpression = append(resultPostfixExpression, operators[len(operators)-1])
		operators = operators[:len(operators)-1]
	}
	return resultPostfixExpression, nil
}
