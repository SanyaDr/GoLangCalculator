package app

import (
	"SecondSprintExam/internal/transport"
	"fmt"
	"net/http"
)

func setServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Sprintln("Приветствую в калькуляторе!")
	})
	http.HandleFunc("/api/v1/calculate", transport.GetExpression)
}
