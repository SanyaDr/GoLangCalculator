package app

import (
	config "SecondSprintExam"
	"context"
	"time"
)

// каждую секунду ожидает получения новой задачии
func setChecker() {
	go func() {
		for {
			ctx, cancel := context.WithCancel(context.Background())
			select {
			case <-ctx.Done():
				cancel()

			case <-time.After(config.CheckNewExpression_Timeout * time.Second):

			}
		}
	}()
}
