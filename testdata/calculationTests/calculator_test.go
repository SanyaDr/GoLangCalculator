package calculationTests

import (
	calculator "GoLangCalculator/pkg/calculation"
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name           string
		giveExpression string
		wantResult     float64
		wantErr        error
	}{
		//simple tests
		{"simple success", "2+2", 4, nil},
		{"simple success", "3*4", 12, nil},
		{"check priority", "2+2*2", 6, nil},
		{"long simple", "2*3+3*2-5*5", -13, nil},
		{"parentheses check", "(3+5)*2", 16, nil},
		{"div by zero", "(2+2)/0", 0, calculator.ErrDivisionByZero},
		{"simple num", "2", 2, nil},
		{"simple parentheses", "(2)+(3)+1", 6, nil},
		{"negative nums", "-5-2-4", -11, nil},
		{"with spaces", "2 + 3 + 6-12 + 1/1 +2 * 5", 10, nil},
		//{"unsupported operation", "2^5", 0, calculator.ErrUnsupportedOperation},
		//more difficult tests
		{"success a lot parentheses", "(1+(4+5+2)-3)+(6+8)", 23, nil},
		{"miss parentheses", "(2+3", 0, calculator.ErrMismatchedParentheses},
		{"invalid expression", "2++2", 0, calculator.ErrInvalidExpression},
		{"success parentheses", "(((2*2)))", 4, nil},
		{"miss parentheses", "(((2+2)*2)))", 0, calculator.ErrMismatchedParentheses},
		{"float expression", "2+2-5+6*34/10-20+5.59", 4.99, nil},
		{"big multiple", "255*34-166*2", 8338, nil},

		{"letters", "3+r-14+ghrkfjd01", 0, calculator.ErrInvalidExpression},
		{"empty expression", "", 0, calculator.ErrEmptyExpression},
		{"tmp1", "(2+2)/8", 0.5, nil},
		{"38", "534%432&23!21", 0, calculator.ErrInvalidExpression},
		{"16", "2++2", 0, calculator.ErrInvalidExpression},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, gotErr := calculator.Calc(tt.giveExpression)
			gotResult, err := strconv.ParseFloat(fmt.Sprintf("%.2f", gotResult), 64)
			if err != nil {
				t.Errorf("internal error! ParseFloat(), got err: %v", err)
			}
			if !errors.Is(gotErr, tt.wantErr) {
				t.Errorf("ERROR: Calc('%v') -> gotError = %v, but wantErr %v", tt.giveExpression, gotErr, tt.wantErr)
			}
			if gotResult != tt.wantResult {
				t.Errorf("ERROR: Calc('%v') -> gotResult = %v, but want %v", tt.giveExpression, gotResult, tt.wantResult)
			}
		})
	}
}
