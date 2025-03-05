package transport

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
	_, exists := expressionStorage[nextId]
	for exists != false {
		_, exists = expressionStorage[nextId]
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

// Получить карту всех выражений
func GetMapAllExpressions() map[int]Expression {
	return expressionStorage
}

func GetUnresolvedOne() (Expression, bool) {
	exists := false
	var expr Expression
	for _, expression := range unresolvedExpressions {
		exists = true
		expr = expression
		break
	}
	if exists {
		delFromUnresolved(expr.Id)
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
	//delFromUnresolved(id)
}
