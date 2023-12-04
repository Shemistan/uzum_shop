package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"group.id":          "my-consumer-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println("failed to creat consumer")
		return
	}

	topics := []string{"my_topic"}
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("failed to send topics")
		return
	}

	log.Println("жду сообщений--- ")

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Println(string(msg.Value))
		} else {
			fmt.Println("error readmsg")
		}
	}
}
