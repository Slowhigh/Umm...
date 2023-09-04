package kafka_client

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/slowhigh/Umm/pkg/logger"
)

type Producer interface {
	PublishMessage(ctx context.Context, msgs ...kafka.Message)
	Close() error
}

type producer struct {
	log     logger.Logger
	brokers []string
	w       *kafka.Writer
}

func NewProducer(log logger.Logger, brokers []string) *producer {
	return &producer{log: log, brokers: brokers, w: NewWriter(brokers, kafka.LoggerFunc(log.Errorf))}
}

func (p *producer) PublishMessage(ctx context.Context, msgs ...kafka.Message) error {
	if err := p.w.WriteMessages(ctx, msgs...); err != nil {
		return err
	}
	return nil
}

func (p *producer) Close() error {
	return p.w.Close()
}
