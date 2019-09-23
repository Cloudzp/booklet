package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	consumer()
}

func consumer() {
	client, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "172.16.65.220:9092",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := client.Subscribe("cloudzp", nil); err != nil {
		fmt.Println(err)
		return
	}

	for {
		msg, err := client.ReadMessage(-1)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("msg", msg)
		}
	}
	client.Close()

}
