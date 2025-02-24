package transport

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type answerResponse struct {
	Result string `json:"result"`
}
type errorResponse struct {
	Error string `json:"error"`
}

// Вывод на экран ответ в виде JSON
func returnAnswer(w http.ResponseWriter, text string) {
	resp := answerResponse{Result: text}
	rData, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = fmt.Fprintf(w, string(rData))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
