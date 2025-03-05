package transport

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func PostStatusCalc(id int, status ExpressionStatus, answer float64) {
	myReq := Expression{
		Id:     id,
		Status: status,
		Result: answer,
	}
	data, err := json.Marshal(myReq)
	if err != nil {
		log.Printf("ERROR: PostStatusCalc(%v) -> json.Marshal() -> Got an error: %v", id, err)
		return
	}
	_, err = http.Post("http://localhost:8080/internal/task", "encoding/json", bytes.NewBuffer(data))
	if err != nil {
		log.Printf("ERROR: PostStatusCalc(%v) -> http.Post() -> Got an error: %v", id, err)
	}
}
