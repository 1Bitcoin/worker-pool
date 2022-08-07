package main

import (
	"fmt"
	"time"
	"worker-pool/workerpool"
)

func main() {
	var tasks []*workerpool.Task

	for i := 1; i <= 100; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, i)
		tasks = append(tasks, task)
	}

	pool := workerpool.NewPool(tasks, 5)
	pool.Run()
}
