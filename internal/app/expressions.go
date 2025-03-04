package app

import "log"

var (
	nextId = 0
)

// Статус решения выражения
type ExpressionStatus string

const (
	Solving ExpressionStatus = "Solving" // Выражение вычисляется
	Failed  ExpressionStatus = "Failed"  // Ошибка вычисления
	Success ExpressionStatus = "Success" //	Вычисление завершено успешно
)

type Expression struct {
	Id         int              `json:"id"`
	Expression string           `json:"expression"`
	Status     ExpressionStatus `json:"status"`
	Result     float64          `json:"result"`
}

var expressionStorage = make(map[int]Expression)     // Список содержащий еще не решенные задачи; По мере решения задачи от сюда удаляются
var unresolvedExpressions = make(map[int]Expression) // Список содержащий абсолютно все задачи (в том числе еще не решенные. Они дублируются и обновляются по id)

func AddNewExpression(expression string) int {
	exp := Expression{
		Id:         nextId,
		Expression: expression,
		Status:     Solving,
		Result:     0.0,
	}
	for exists := true; exists != false; _, exists = expressionStorage[nextId] {
		if exists {
			log.Printf("ERROR: Expression already exists!")
		}
		nextId++
	}
	unresolvedExpressions[nextId] = exp
	expressionStorage[nextId] = exp
	nextId++
	return nextId - 1
}

func delFromUnresolved(id int) {
	_, exists := unresolvedExpressions[id]
	if !exists {
		log.Printf("ERROR: delFromUnresolved(%v) -> unresolved Expression does not exist!", id)
	}
	delete(unresolvedExpressions, id)
}

// TODO удалить лишние функции
// Обновить статус выражение на Failed
func ExpressionFailure(id int) {
	expr, exists := expressionStorage[id]
	if !exists {
		log.Printf("ERROR: ExpressionFailure(%v) -> Expression does not exist!", id)
		return
	}
	expr.Status = Failed
	expressionStorage[id] = expr
	delFromUnresolved(id)
}

// Обновить статус выражения на Success
func ExpressionSuccess(id int, result float64) {
	expr, exists := expressionStorage[id]
	if !exists {
		log.Printf("ERROR: ExpressionSuccess(%v) -> Expression does not exist!", id)
	}
	expr.Result = result
	expr.Status = Success
	expressionStorage[id] = expr
	delFromUnresolved(id)
}

// Получить карту всех выражений
func GetMapAllExpressions() map[int]Expression {
	return expressionStorage
	//allExps := make([]Expression, 0, len(expressionStorage))
	//for _, Expression := range expressionStorage {
	//	allExps = append(allExps, Expression)
	//}
	//return allExps
}

func GetUnresolvedOne() (Expression, bool) {
	exists := false
	var expr Expression
	for _, expression := range unresolvedExpressions {
		exists = true
		expr = expression
		break
	}
	return expr, exists
}

func UpdateExpressionStatus(id int, status ExpressionStatus, result float64) {
	expr, exists := expressionStorage[id]
	if !exists {
		log.Printf("ERROR: UpdateExpressionStatus(%v) -> Expression does not exist!", id)
	}
	expr.Status = status
	if status == Success {
		expr.Result = result
	} else {
		expr.Result = 0
	}
	expressionStorage[id] = expr
	delFromUnresolved(id)
}
