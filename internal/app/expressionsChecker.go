package app

import (
	"SecondSprintExam/config"
	"SecondSprintExam/internal/transport"
	calculator "SecondSprintExam/pkg/calculation"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

var threadsCount = 0

// Каждую секунду ожидает получения новой задачи
func setChecker() {
	go func() {
		for {
			ctx, cancel := context.WithCancel(context.Background())
			select {
			case <-ctx.Done():
				cancel()

			case <-time.After(config.CheckNewExpression_Timeout * time.Millisecond):
				if !transport.GetExistUnresolved() {
					break
				}
				url := "http://localhost:8080/internal/task"

				myReq, err := http.NewRequest("GET", url, nil)
				if err != nil {
					log.Printf("ERROR: setChecker() -> http.NewRequest err: %v", err)
					break
				}
				myReq.Header.Set("FromWhat", "FromCalc")
				client := &http.Client{}
				resp, err := client.Do(myReq)
				if err != nil {
					log.Printf("ERROR: setChecker() -> client.Do err: %v", err)
				}

				//resp, err := http.Get(url)
				//if err != nil {
				//	log.Printf("ERROR: SetChecker() -> http.Get(%v): %v", url, err)
				//	break
				//}
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Printf("ERROR: SetChecker() -> io.ReadAll(): %v", err)
				}
				var expr transport.TaskResponse
				err = json.Unmarshal(body, &expr)
				if err != nil {
					log.Printf("ERROR: SetChecker() -> json.Unmarshal(): %v", err)
				}

				for threadsCount >= config.GetComputingPower() {
					time.After(250 * time.Millisecond)
				}
				go func() {
					mu.Lock()
					threadsCount++
					mu.Unlock()

					defer func() {
						mu.Lock()
						threadsCount--
						mu.Unlock()
					}()

					ans, err := calculator.Calc(expr.Expression)
					if err != nil {
						transport.PostStatusCalc(expr.Id, transport.Failed, 0)
						return
					}
					transport.PostStatusCalc(expr.Id, transport.Success, ans)
					log.Printf("Задача %v решена", expr.Id)
				}()
			}
		}
	}()
}
