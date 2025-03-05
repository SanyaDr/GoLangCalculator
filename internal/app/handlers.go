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
	http.HandleFunc("/api/v1/calculate", transport.GetExpressionHandler)
	// "/api/v1/expressions" - получение списка выражений
	// "/api/v1/expressions/:id" - Получение выражения по его идентификатору
	http.HandleFunc("/api/v1/expressions", transport.GetAllExpressionsHandler)

	//Меняем суть, теперь оно отдает не таски и целиком выражения
	// "/internal/task" - Получение задачи для выполнения
	http.HandleFunc("/internal/task", transport.GetTaskHandler)
	// "/internal/task" - Прием результата обработки данных
	// ?????????

}
