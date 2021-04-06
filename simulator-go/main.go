package main

import (
	"fmt"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafkaProduce "github.com/souzarodrigo61/simulator-go/applications/kafka"
	"github.com/souzarodrigo61/simulator-go/infra/kafka"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *confluentKafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)

	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafkaProduce.Produce(msg)
	}
}