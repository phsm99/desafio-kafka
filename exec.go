package main

import (
	c "desafio/consumer"
	p "desafio/producer"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan ckafka.Event)
	producer := p.NewKafkaProducer()

	p.Publish("Oi", "mensageiro_tpc", producer, deliveryChan)
	go p.DeliveryReport(deliveryChan)

	kafkaProcessor := c.NewKafkaProcessor(producer, deliveryChan)
	kafkaProcessor.Consume()
}
