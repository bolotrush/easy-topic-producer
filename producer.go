package easy_topic_producer

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

//TopicProducer adapter for sending messages to kafka topic
type TopicProducer interface {
	SendMessage(key, value []byte) error
}

type producer struct {
	syncProducer sarama.SyncProducer
	topic        string
}

//NewTopicProducer instance
func NewTopicProducer(brokers []string, topic string) (TopicProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_4_0_0

	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Retry.Max = 10
	cfg.Producer.RequiredAcks = sarama.WaitForLocal

	p, err := sarama.NewSyncProducer(
		brokers,
		cfg)
	if err != nil {
		return nil, fmt.Errorf("NewTopicProducer: %v", err)
	}

	return &producer{
		syncProducer: p,
		topic:        topic,
	}, nil
}

//SendMessage to kafka
func (p *producer) SendMessage(key, value []byte) error {
	_, _, err := p.syncProducer.SendMessage(&sarama.ProducerMessage{
		Topic:     p.topic,
		Key:       sarama.ByteEncoder(key),
		Value:     sarama.ByteEncoder(value),
		Timestamp: time.Now().UTC(),
	})
	return err
}
