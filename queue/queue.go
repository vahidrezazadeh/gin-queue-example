package queue

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Queue struct {
	client *redis.Client
	queue  string
}

func NewQueue(client *redis.Client, queueName string) *Queue {
	return &Queue{
		client: client,
		queue:  queueName,
	}
}

func (q *Queue) AddToQueue(item string) error {
	return q.client.LPush(context.Background(), q.queue, item).Err()
}

func (q *Queue) DoQueue() {
	for {
		item, err := q.client.RPop(context.Background(), q.queue).Result()
		if err != nil {
			break
		}
		// process the item
		q.processItem(item)
	}
}

func (q *Queue) processItem(item string) {
	// todo every tasks
	fmt.Println("Processing item:", item)
}
