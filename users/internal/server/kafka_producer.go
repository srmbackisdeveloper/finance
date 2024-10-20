package server

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func NewKafkaProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{os.Getenv("KAFKA_BROKER")}
	return sarama.NewSyncProducer(brokers, config)
}

func (s *Server) SendUserRegistered(email, code string) error {
	topic := "user-registered"

	message := struct {
		Email string `json:"email"`
		Code  string `json:"code"`
	}{
		Email: email,
		Code:  code,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to serialize message: %v", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(messageBytes),
	}

	partition, offset, err := s.producer.SendMessage(msg)
	if err != nil {
		return fmt.Errorf("failed to send Kafka message: %v", err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
