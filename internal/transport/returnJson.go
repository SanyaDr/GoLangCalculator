package transport

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type errorResponse struct {
	Error string `json:"error"`
}
type newExpressionRequest struct {
	Id int `json:"id"`
}
type TaskResponse struct {
	Id         int    `json:"id"`
	Expression string `json:"expression"`
}
type ExpressionResponse struct {
	Id     int              `json:"id"`
	Status ExpressionStatus `json:"status"`
	Result float64          `json:"result"`
}
type AllExpressionsResponse struct {
	Expressions []ExpressionResponse `json:"expressions"`
}
type ExpressionByIdResponse struct {
	Expression ExpressionResponse `json:"expression"`
}

// Вывод на экран ошибки в формате JSON
func returnError(w http.ResponseWriter, text string, statusCode int) {
	resp := errorResponse{Error: text}
	rData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Error(w, string(rData), statusCode)
}

// Вывод на экран ответа по эндпоинту '/api/v1/calculate'
func returnSuccessAddingNewExpression(w http.ResponseWriter, id int) {
	resp := newExpressionRequest{Id: id}
	rData, err := json.Marshal(resp)
	if err != nil {
		returnError(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = fmt.Fprintf(w, string(rData))
	if err != nil {
		returnError(w, err.Error(), http.StatusInternalServerError)
	}
}

// GetAllExpressionsHandler Вывод на экран всех выражений
func GetAllExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	var rData []byte
	var err error

	// Получаем id выражения если есть
	query := r.URL.Query()
	var selectedId string
	for key, values := range query {
		if strings.EqualFold(key, "id") {
			selectedId = values[0]
			break
		}
	}
	// Если iD не задано, выводим все выражения
	if selectedId == "" {
		allExpressions := GetAllExpressions()
		expressionResponses := make([]ExpressionResponse, 0, len(allExpressions))
		for _, exp := range allExpressions {
			expressionResponses = append(expressionResponses, ExpressionResponse{exp.Id, exp.Status, exp.Result})
		}
		resp := AllExpressionsResponse{expressionResponses}
		rData, err = json.Marshal(resp)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		// Получить выражение по iD
		allMapExpressions := GetMapAllExpressions()
		curId, err := strconv.Atoi(selectedId)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		expression, exists := allMapExpressions[curId]
		exprById := ExpressionResponse{expression.Id, expression.Status, expression.Result}
		resp := ExpressionByIdResponse{exprById}
		if !exists {
			returnError(w, "Does not exist", http.StatusNotFound)
			return
		}
		rData, err = json.Marshal(resp)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	_, err = fmt.Fprintf(w, string(rData))
	if err != nil {
		returnError(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.Header.Get("FromWhat") == "FromCalc" {
		newExpression, exists := GetUnresolvedOne()
		if !exists {
			returnError(w, "No tasks", http.StatusInternalServerError)
			return
		}
		resp := TaskResponse{
			Id:         newExpression.Id,
			Expression: newExpression.Expression,
		}
		rData, err := json.Marshal(resp)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = fmt.Fprintf(w, string(rData))
	} else if r.Method == "POST" {
		var myReq Expression

		body, err := io.ReadAll(r.Body)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(body, &myReq)
		UpdateExpressionStatus(myReq.Id, myReq.Status, myReq.Result)
		// Обновить статус выражения

	} else {
		returnError(w, "Method not allowed", http.StatusUnprocessableEntity)
		return
	}
}
