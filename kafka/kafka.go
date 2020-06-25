package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	brokerAddrs := []string{"<kafka-brokers>:9092"}
	config := sarama.NewConfig()
	config.Version = sarama.V0_10_2_1
	admin, err := sarama.NewClusterAdmin(brokerAddrs, config)
	if err != nil {
		log.Fatal("Error while creating cluster admin: ", err.Error())
	}
	defer admin.Close()

	err = admin.CreateTopic("gaurav", &sarama.TopicDetail{
		NumPartitions:     1,
		ReplicationFactor: 1,
	}, false)
	if err != nil {
		log.Fatal("Error while creating topic: ", err.Error())
	}
}
