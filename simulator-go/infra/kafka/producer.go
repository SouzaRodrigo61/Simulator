package kafka

import (
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
)

// NewKafkaProducer creates a ready to go kafka.Producer instance
func NewKafkaProducer() *confluentKafka.Producer {
	configMap := &confluentKafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KafkaBootstrapServers"),
	}
	p, err := confluentKafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

// Publish is simple function created to publish new message to kafka
func Publish(msg string, topic string, producer *confluentKafka.Producer) error {
	message := &confluentKafka.Message{
		TopicPartition: confluentKafka.TopicPartition{Topic: &topic, Partition: confluentKafka.PartitionAny},
		Value:          []byte(msg),
	}
	err := producer.Produce(message, nil)
	if err != nil {
		return err
	}
	return nil
}