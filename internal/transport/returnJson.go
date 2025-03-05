package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// TODO Del unused
type answerResponse struct {
	Result string `json:"result"`
}
type errorResponse struct {
	Error string `json:"error"`
}
type newExpressionRequest struct {
	Id int `json:"id"`
}
type taskResponse struct {
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

//TODO Del com
//// Вывод на экран ответ в виде JSON
//func returnAnswer(w http.ResponseWriter, text string) {
//	resp := answerResponse{Result: text}
//	rData, err := json.Marshal(resp)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//	_, err = fmt.Fprintf(w, string(rData))
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}

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

// TODO Test this shit!
// TODO Не забудь протестировать каджую строку особенно получение по id и поведение его exist
// Вывод на экран всех выражений
func GetAllExpressionsHandler(w http.ResponseWriter, r *http.Request) {
	var rData []byte
	var err error

	//selectedId := r.URL.Query().Get("id")
	query := r.URL.Query()
	var selectedId string
	for key, values := range query {
		if strings.EqualFold(key, "id") {
			selectedId = values[0] // Берем первое значение
			break
		}
	}
	// Получить все выражения
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

// TODO Test This Shit -> TTS
// TODO пользователь может сам обратиться к этой хуйне и невольно удалить еще не решенное выражение. Добавь флаг обозначающий обращение именно самого калькулятора
// TODO потестируй еще раз без удаления нерешенного сразу после запроса. Придумай более элегантный способ
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		newExpression, exists := GetUnresolvedOne()
		if !exists {
			returnError(w, "No tasks", http.StatusInternalServerError)
			return
		}
		resp := taskResponse{
			Id:         newExpression.Id,
			Expression: newExpression.Expression,
		}
		rData, err := json.Marshal(resp)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = fmt.Fprintf(w, string(rData))
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else if r.Method == "POST" {
		var myReq Expression

		body, err := ioutil.ReadAll(r.Body)
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
