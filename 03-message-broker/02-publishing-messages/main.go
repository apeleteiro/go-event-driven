package main

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-redisstream/pkg/redisstream"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/redis/go-redis/v9"
	"os"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"),
	})

	publisher, err := redisstream.NewPublisher(redisstream.PublisherConfig{
		Client: rdb,
	}, logger)
	if err != nil {
		panic(err)
	}

	msg1 := message.NewMessage(watermill.NewUUID(), []byte("50"))
	msg2 := message.NewMessage(watermill.NewUUID(), []byte("100"))

	err = publisher.Publish("progress", msg1)
	if err != nil {
		panic(err)
	}

	err = publisher.Publish("progress", msg2)
	if err != nil {
		panic(err)
	}
}
