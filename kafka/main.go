package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	msg, err := rdb.Get(context.Background(), "first_key").Result()
	if err != nil {
		fmt.Println(err)
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP("kafka:9093"),
		Topic:                  "test_topic",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte(msg),
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
