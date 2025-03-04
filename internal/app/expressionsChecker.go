package app

import (
	"SecondSprintExam/config"
	"SecondSprintExam/internal/transport"
	calculator "SecondSprintExam/pkg/calculation"
	"context"
	"log"
	"time"
)

var threadsCount = 0

// каждую секунду ожидает получения новой задачи
func setChecker() {
	go func() {
		for {
			ctx, cancel := context.WithCancel(context.Background())
			select {
			case <-ctx.Done():
				cancel()

			case <-time.After(config.CheckNewExpression_Timeout * time.Second):
				// TODO тут не обращаемся напрямую а через http
				expr, exists := transport.GetUnresolvedOne()
				if !exists {
					break
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
