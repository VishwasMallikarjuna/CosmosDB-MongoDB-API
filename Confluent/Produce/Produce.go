package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	//p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": os.Args[1]})

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":                     os.Args[1],
		"sasl.mechanisms":                       "PLAIN",
		"security.protocol":                     "sasl_ssl",
		"sasl.username":                         "client",
		"sasl.password":                         "client-secret",
		"ssl.ca.location":                       "C:/kafka/bundle/client.cer.pem",
		"ssl.endpoint.identification.algorithm": "https"})

	fmt.Println(p.Events())

	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	if err != nil {
		panic(err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Message Produced")
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := os.Args[2]
	for _, word := range []string{"Produced Message from Go kafka client Confluent"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(30 * 1000)

	//NewHealthChecker()
}

func NewHealthChecker() {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers":                     os.Args[1],
		"sasl.mechanisms":                       "PLAIN",
		"security.protocol":                     "sasl_ssl",
		"sasl.username":                         "kafka",
		"sasl.password":                         "kafka-secret",
		"ssl.endpoint.identification.algorithm": "https"}

	client, err := kafka.NewAdminClient(kafkaConfig)

	if err != nil {
		fmt.Println("error constructing Kafka admin client: %w", err)
	}

	fmt.Println("client details ", client)

}
