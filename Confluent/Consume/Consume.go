package main

/**
 * Copyright 2020 Confluent Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	// Create Consumer instance
	// c, err := kafka.NewConsumer(&kafka.ConfigMap{
	// 	"bootstrap.servers": os.Args[1],
	// 	"group.id":          "go_example_group_1",
	// 	"auto.offset.reset": "earliest",
	// })

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":                     os.Args[1],
		"group.id":                              "go_example_group_1",
		"auto.offset.reset":                     "earliest",
		"sasl.mechanisms":                       "PLAIN",
		"security.protocol":                     "sasl_ssl",
		"sasl.username":                         "client",
		"sasl.password":                         "client-secret",
		"ssl.ca.location":                       "C:/kafka/bundle/client.cer.pem",
		"ssl.endpoint.identification.algorithm": "https"})
	if err != nil {
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	// Subscribe to topic
	c.SubscribeTopics([]string{os.Args[2]}, nil)
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	//totalCount := 0
	run := true
	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			msg, err := c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			//recordKey := string(msg.Key)
			recordValue := msg.Value
			if err != nil {
				fmt.Printf("Failed to decode JSON at offset %d: %v", msg.TopicPartition.Offset, err)
				continue
			}
			fmt.Printf("Consumed record with value:%s\n", recordValue)
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()

}
