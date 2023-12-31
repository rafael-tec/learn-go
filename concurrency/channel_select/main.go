package main

import (
	"fmt"
	"time"
)

type Message struct {
	id   int
	Body string
}

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	// It could be RabbitMQ
	go func() {
		time.Sleep(time.Second * 3)
		ch1 <- Message{id: 12345, Body: "Message from RabbitMQ"}
	}()

	// It could be Kafka
	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- Message{id: 98765, Body: "Message from Kafka"}
	}()

	for {
		fmt.Println("One more iterate")
		select {
		case msgByRabbit := <-ch1:
			fmt.Printf(
				"Received message from RabbitMQ\n[ID: %d - Body: %s]\n",
				msgByRabbit.id,
				msgByRabbit.Body,
			)
		case msgByKafka := <-ch2:
			fmt.Printf(
				"Received message from Kafka\n[ID: %d - Body: %s]\n",
				msgByKafka.id,
				msgByKafka.Body,
			)
		case <-time.After(time.Second * 10):
			fmt.Printf("Timed out after 3 seconds waiting\n")
		}
	}
}
