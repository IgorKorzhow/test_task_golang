package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// Creat crone instance
	c := cron.New()

	// Added tasks
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("Task executed at:", time.Now())
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}

	// Run scheduler
	c.Start()

	// Stop scheduler when stopping program
	defer c.Stop()

	// lock main thread for block complete programme
	select {}
}
