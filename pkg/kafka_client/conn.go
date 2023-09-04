package kafka_client

import (
	"context"

	"github.com/segmentio/kafka-go"
	"github.com/slowhigh/Umm/pkg/constants"
)

func NewKafkaConn(ctx context.Context, kafkaCfg *Config) (*kafka.Conn, error) {
	return kafka.DialContext(ctx, constants.TCP, kafkaCfg.Brokers[0])
}