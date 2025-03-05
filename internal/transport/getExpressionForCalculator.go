package transport

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type MyRequest struct {
	Expression string `json:"expression"`
}

// GetExpression Метод чтения выражения и отправка его в Calculation
func GetExpression(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var myReq MyRequest
	// Проверяем что получен именно POST метод
	if r.Method != http.MethodPost {
		log.Println("ERROR: получен не POST метод!")
		returnError(w, "Method is not allowed", http.StatusUnprocessableEntity)
		return
	}

	// Получаем тело запроса и сам expression
	body, err := io.ReadAll(r.Body)

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

	curId := AddNewExpression(myReq.Expression)
	log.Printf("Получено новое выражение: %v", myReq.Expression)
	returnSuccessAddingNewExpression(w, curId)
}
