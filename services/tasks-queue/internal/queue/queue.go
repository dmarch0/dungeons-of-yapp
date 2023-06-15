package queue

import (
	"github.com/go-co-op/gocron"
	"time"
	"fmt"
)

func StartQueue() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(60).Seconds().Do(func() {
		fmt.Println("5 seconds passed")
	})

	s.StartBlocking();
}