package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

func main() {

	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Println("failed to creat producer")
		return
	}

	producer.Flush(15 * 1000)
	topic := "my_topic"

	i := 0
	for {
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(fmt.Sprintf("Message number %d", i)),
		}

		err := producer.Produce(msg, nil)
		if err != nil {
			fmt.Println("failed to send message")
			return
		}

		log.Println("Sent message number is", i)
		i++
		time.Sleep(2 * time.Second)

	}
}
