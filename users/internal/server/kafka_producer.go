package server

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func (s *Server) KafkaProducer(topic string, message string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // Ensure all brokers acknowledge
	config.Producer.Retry.Max = 5                    // Retry up to 5 times
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		log.Fatalf("failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalf("failed to send Kafka message: %v", err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
