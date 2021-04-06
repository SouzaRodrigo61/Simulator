package kafka

import (
	"fmt"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

// MARK: - Typing

// Consumer holds all consumer logic and settings of Apache Kafka connections/
// Also has a Message channel which is a channel where the messages are going to be pushed
type Consumer struct {
	MsgChan chan *confluentKafka.Message
}

// MARK: - Method

// NewKafkaConsumer creates a new KafkaConsumer struct with its message channel as dependency
func NewKafkaConsumer(msgChan chan *confluentKafka.Message) *Consumer {
	return &Consumer{
		MsgChan: msgChan,
	}
}

// Consume consumes all message pulled from apache kafka and sent it to message channel
func (k *Consumer) Consume() {
	configMap := &confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
		"group.id":          os.Getenv("KafkaConsumerGroupId"),
	}

	c, err := confluentKafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("error consuming kafka message:" + err.Error())
	}

	topics := []string{os.Getenv("KafkaReadTopic")}

	err = c.SubscribeTopics(topics, nil)
	if err != nil {
		log.Fatalf("Erro para subscrever nos topicos do kafka", err.Error())
	}

	fmt.Println("Kafka consumer has been started")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			k.MsgChan <- msg
		}
	}
}