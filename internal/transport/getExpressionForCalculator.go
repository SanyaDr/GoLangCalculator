package transport

import (
	"SecondSprintExam/pkg/calculation"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type MyRequest struct {
	Expression string `json:"expression"`
}

// GetExpression Метод чтения выражения и отправка его в Calculation
func GetExpression(w http.ResponseWriter, r *http.Request) {
	var myReq MyRequest
	// Проверяем что получен именно POST метод
	if r.Method != http.MethodPost {
		log.Println("ERROR: получен не POST метод!")
		returnError(w, "Method is not allowed", http.StatusUnprocessableEntity)
		return
	}

	// Получаем тело запроса и сам expression
	body, err := ioutil.ReadAll(r.Body)

	//log.Printf("body:\n %v", string(body))
	//log.Printf("Request.Header.content-type: %v", r.Header.Get("Content-Type"))

	if err != nil {
		log.Printf("ERROR: Ошибка получения данных запроса! Текст ошибки: %v", err)
		returnError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	err = json.Unmarshal(body, &myReq)
	if err != nil {
		log.Printf("ERROR: Ошибка обработки JSON! Текст ошибки: %v", err)
		returnError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем полученный expression в Calculator
	result, err := calculator.Calc(myReq.Expression)
	if err != nil {
		returnError(w, err.Error(), http.StatusUnprocessableEntity)
		log.Printf("ERROR: Ошибка при вычислении! Текст ошибки: %v", err)
		return
	}

	returnAnswer(w, fmt.Sprint(result))
	log.Printf("Получен результат: %v", result)

}
