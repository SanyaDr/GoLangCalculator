package app

import (
	"GoLangCalculator/config"
	"log"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex

func RunServer() {
	setEndpoints()
	setChecker()

	log.Printf("Сервер запущен на http://localhost:%v", config.DefaultLaunchPort)
	log.Printf("Время запуска: %v\n", time.Now().Format("15:04:05"))

	if err := http.ListenAndServe(":"+config.DefaultLaunchPort, nil); err != nil {
		log.Fatalf("Ошибка запуска сервера! Текст ошибки: \n%v", err)
	}

}
