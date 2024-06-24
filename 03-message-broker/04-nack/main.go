package main

import (
	"context"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
)

type AlarmClient interface {
	StartAlarm() error
	StopAlarm() error
}

func ConsumeMessages(sub message.Subscriber, alarmClient AlarmClient) {
	messages, err := sub.Subscribe(context.Background(), "smoke_sensor")
	if err != nil {
		panic(err)
	}

	for msg := range messages {
		smokeDetected := string(msg.Payload)
		if smokeDetected == "1" {
			err = alarmClient.StartAlarm()
			if err != nil {
				fmt.Println("Error starting alarm:", err)
				msg.Nack()
				continue
			}
		} else {
			err = alarmClient.StopAlarm()
			if err != nil {
				fmt.Println("Error stopping alarm:", err)
				msg.Nack()
				continue
			}

		}
		msg.Ack()
	}
}
