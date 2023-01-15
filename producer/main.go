package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

func main() {

	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "kafka", "acks": "1"})
	if err != nil {
		panic(err)
	}

	defer p.Close()

	topic := "kafka-test"
	dChan := make(chan kafka.Event)
	go pDelivery(dChan)
	var messageCount = 1
	for {
		time.Sleep(3 * time.Second)
		fmt.Printf("Producing message %d.. \n", messageCount)
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte("Testing"),
			Key:            []byte("key-test"),
		}, dChan)
		messageCount++
	}

	//p.Flush(15 * 1000)
}

func pDelivery(dChan <-chan kafka.Event) {
	for e := range dChan {
		switch ev := e.(type) {
		case *kafka.Message:
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed, error: %s \n", ev.TopicPartition.Error.Error())
			} else {
				fmt.Printf("Delivered message: %s | TopicPartition: %s \n", string(ev.Value), ev.TopicPartition)
			}
		}
	}
}
