package consumer

import (
	"fmt"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaProcessor struct {
	Producer     *ckafka.Producer
	DeliveryChan chan ckafka.Event
}

func NewKafkaProcessor(producer *ckafka.Producer, deliveryChan chan ckafka.Event) *KafkaProcessor {
	return &KafkaProcessor{
		Producer:     producer,
		DeliveryChan: deliveryChan,
	}
}

func (k *KafkaProcessor) Consume() {
	configMap := &ckafka.ConfigMap{
		"bootstrap.servers": "host.docker.internal:9094",
		"group.id":          "mensageiro_grp",
		"auto.offset.reset": "earliest",
	}
	c, err := ckafka.NewConsumer(configMap)

	if err != nil {
		panic(err)
	}

	topics := []string{"mensageiro_tpc"}
	c.SubscribeTopics(topics, nil)

	fmt.Println("kafka consumer has been started")
	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			fmt.Println("Mensagem:", string(msg.Value))
		}
	}
}
