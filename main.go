package main

import (
	"github.com/redis/go-redis/v9"
	queue2 "github.com/vahidrezazadeh/gin-queue-example/queue"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	queue := queue2.NewQueue(client, "my_queue")

	// add items to the queue
	queue.AddToQueue("item1")
	queue.AddToQueue("item2")
	queue.AddToQueue("item3")

	// start processing the queuej
	queue.DoQueue()
}
