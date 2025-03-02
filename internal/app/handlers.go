package app

import (
	"SecondSprintExam/internal/transport"
	"fmt"
	"net/http"
)

// setEndpoints - установка endpoint'ов
func setEndpoints() {
	// "/" - Главная страница калькулятора
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Приветствую в калькуляторе!")
	})
	// "/api/v1/calculate" - точка принятия новых выражений
	http.HandleFunc("/api/v1/calculate", transport.GetExpression)
	// "/api/v1/expressions" - получение списка выражений
	// http.handleFunc()
	// "/api/v1/expressions/:id" - Получение выражения по его идентификатору
	// http.handleFunc()
	// "/internal/task" - Получение задачи для выполнения
	// http.handleFunc()
	// "/internal/task" - Прием результата обработки данных
	// ?????????

}
