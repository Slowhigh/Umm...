package kafka_client

import (
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

func NewWriter(brokers []string, errLogger kafka.Logger) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(brokers...),
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireAll,
		MaxAttempts:  writerMaxAttempts,
		ErrorLogger:  errLogger,
		Compression:  compress.Snappy,
		ReadTimeout:  writerReadTimeout,
		WriteTimeout: writerWriteTimeout,
		BatchTimeout: batchTimeout,
		BatchSize:    batchSize,
		Async:        false,
	}
}
