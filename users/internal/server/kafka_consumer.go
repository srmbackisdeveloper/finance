package server

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

func (s *Server) KafkaConsumer(topic string, groupID string) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, groupID, config)
	if err != nil {
		log.Fatalf("failed to create Kafka consumer group: %v", err)
	}
	defer consumerGroup.Close()

	consumer := Consumer{}

	for {
		err := consumerGroup.Consume(context.Background(), []string{topic}, &consumer)
		if err != nil {
			log.Fatalf("error consuming Kafka topic: %v", err)
		}
	}
}

type Consumer struct{}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim processes Kafka messages
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("Message received: key=%s, value=%s, topic=%s, partition=%d, offset=%d\n",
			string(msg.Key), string(msg.Value), msg.Topic, msg.Partition, msg.Offset)

		// Mark message as processed
		session.MarkMessage(msg, "")
	}
	return nil
}
