package processor

import (
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type Processor struct{}

func NewProcessor() Processor {
	return Processor{}
}

func (p *Processor) Process(d rabbitmq.Delivery) rabbitmq.Action {
	logrus.Info(string(d.Body))
	logrus.Info(string(d.RoutingKey))
	return rabbitmq.Ack
}
