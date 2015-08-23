package main

import (
	"github.com/jrallison/go-workers"
)

const (
	defaultQueue = "default"
)

func main() {
	workers.Configure(map[string]string{
		"server":    "localhost:6379",
		"process":   "test-go-workers",
		"namespace": "goworkers-sidekiq",
	})
	workers.EnqueueIn(defaultQueue, "HelloWorker", 1.0, []string{"Soulou"})
}
