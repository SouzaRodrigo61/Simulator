package kafka

import (
	"encoding/json"
	confluentKafka "github.com/confluentinc/confluent-kafka-go/kafka"
	route "github.com/souzarodrigo61/simulator-go/applications/router"
	"github.com/souzarodrigo61/simulator-go/infra/kafka"
	"log"
	"os"
	"time"
)

// Produce is responsible to publish the positions of each request
// Example of a json request:
//{"clientId":"1","routerId":"1"}
//{"clientId":"2","routerId":"2"}
//{"clientId":"3","routerId":"3"}
func Produce(msg *confluentKafka.Message) {
	producer := kafka.NewKafkaProducer()
	newRoute := route.NewRoute()
	err := json.Unmarshal(msg.Value, &newRoute)
	if err != nil {
		log.Println("Erro ao Unmarshal: " + err.Error())
	}
	err = newRoute.LoadPositions()
	if err != nil {
		log.Println("Erro ao carregar o load positions: " + err.Error())
	}
	positions, err := newRoute.ExportJsonPositions()
	if err != nil {
		log.Println("Erro ao ExportJsonPositions: " + err.Error())
	}
	for _, p := range positions {
		err = kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		if err != nil {
			log.Println("Erro ao Publish: " + err.Error())
		}
		time.Sleep(time.Millisecond * 500)
	}
}