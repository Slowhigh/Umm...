package app

import (
	"context"
	"net"
	"strconv"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/slowhigh/Umm/internal/domain"
	"github.com/slowhigh/Umm/pkg/constants"
	eventStores "github.com/slowhigh/Umm/pkg/event_stores"
	kafkaClient "github.com/slowhigh/Umm/pkg/kafka_client"
)

func (a *app) connectKafkaBrokers(ctx context.Context) error {
	kafkaConn, err := kafkaClient.NewKafkaConn(ctx, a.cfg.Kafka)
	if err != nil {
		return errors.Wrap(err, "kafka.NewKafkaConn")
	}

	a.kafkaConn = kafkaConn

	brokers, err := kafkaConn.Brokers()
	if err != nil {
		return errors.Wrap(err, "kafkaConn.Brokers")
	}

	a.log.Infof("kafka connected (brokers: %+v)", brokers)
	return nil
}

func (a *app) initKafkaTopics(ctx context.Context) {
	controller, err := a.kafkaConn.Controller()
	if err != nil {
		a.log.Error("kafkaConn.Controller err: %v", err)
		return
	}

	controllerURI := net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port))
	a.log.Infof("kafka controller uri (controllerURI: %s)", controllerURI)

	conn, err := kafka.DialContext(ctx, constants.TCP, controllerURI)
	if err != nil {
		a.log.Errorf("initKafkaTopics.DialContext err:% v", err)
		return
	}
	defer conn.Close()

	a.log.Infof("established new kafka controller connection (controllerURI: %s)", controllerURI)

	bankAccountAggregateTopic := eventStores.GetKafkaAggregateTypeTopic(a.cfg.KafkaPublisherConfig, string(domain.BankAccountAggregateType))

	if err := conn.CreateTopics(bankAccountAggregateTopic); err != nil {
		a.log.WarnErrMsg("kafkaConn.CreateTopics", err)
		return
	}

	//TODO: 왜 topic을 2개를 생성했는지 확인해 볼 것
	if err := conn.CreateTopics(bankAccountAggregateTopic); err != nil {
		a.log.Errorf("kafkaConn.CreateTopics (err: %v)", err)
	}

	a.log.Infof("[kafka topics created or already exists]: %+v", []kafka.TopicConfig{bankAccountAggregateTopic})
}
