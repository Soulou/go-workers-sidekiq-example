package main

import (
	"fmt"

	"github.com/jrallison/go-workers"
)

func main() {
	workers.Configure(map[string]string{
		"server":    "localhost:6379",
		"process":   "test-worker",
		"namespace": "goworkers-sidekiq",
	})

	workers.Process("default", HelloWorker, 5)

	workers.Run()
}

func HelloWorker(args *workers.Msg) {
	fmt.Println("Hello World!")
}
