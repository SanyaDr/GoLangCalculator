package transport

import (
	"SecondSprintExam/internal/app"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

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
// Вывод на экран всех выражений
func GetAllExpressions(w http.ResponseWriter, r *http.Request) {
	expressions := app.GetMapAllExpressions()
	var rData []byte
	var err error

	selectedId := r.URL.Query().Get("id")
	if selectedId == "" {
		rData, err = json.Marshal(expressions)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		curId, err := strconv.Atoi(selectedId)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
		}
		rData, err = json.Marshal(expressions[curId])
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
		}
	}

	_, err = fmt.Fprintf(w, string(rData))
	if err != nil {
		returnError(w, err.Error(), http.StatusInternalServerError)
	}
}

// TODO Test This Shit -> TTS
func GetTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		newExpression, exists := app.GetUnresolvedOne()
		if !exists {
			returnError(w, "No tasks", http.StatusInternalServerError)
		}
		resp := taskResponse{
			Id:         newExpression.Id,
			Expression: newExpression.Expression,
		}
		rData, err := json.Marshal(resp)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = fmt.Fprintf(w, string(rData))
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
		}
	} else if r.Method == "POST" {
		var myReq app.Expression

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			returnError(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(body, &myReq)
		app.UpdateExpressionStatus(myReq.Id, myReq.Status, myReq.Result)
		// Обновить статус выражения

	} else {
		http.Error(w, "Method not allowed", http.StatusUnprocessableEntity)
	}
}
