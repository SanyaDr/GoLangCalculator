package app

import (
	"log"
	"net/http"
	"time"
)

func RunServer() {
	setServer()

	log.Println("Сервер запущен на http://localhost:8080")
	log.Printf("Время запуска: %v\n", time.Now().Format("15:04:05"))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера! Текст ошибки: \n%v", err)
	}

}
