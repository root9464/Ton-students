package kafka_util

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

// CreateKafkaProducer создает и возвращает Kafka Producer
func CreateKafkaProducer(brokers string) (*kafka.Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		log.Printf("Error creating Kafka Producer: %v", err)
		return nil, err
	}

	log.Println("Kafka Producer created successfully")
	return producer, nil
}

// CreateKafkaConsumer создает и возвращает Kafka Consumer
func CreateKafkaConsumer(brokers, groupID string, topics []string) (*kafka.Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		log.Printf("Error creating Kafka Consumer: %v", err)
		return nil, err
	}

	// Подписка на указанные топики
	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Printf("Error subscribing to topics: %v", err)
		return nil, err
	}

	log.Println("Kafka Consumer created and subscribed to topics successfully")
	return consumer, nil
}
