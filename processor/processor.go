package processor

import (
	"encoding/json"

	globalmodel "github.com/GarnBarn/common-go/model"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type Processor struct {
	rabbitmqPublisher *rabbitmq.Publisher
}

func NewProcessor(rabbitmqPublisher *rabbitmq.Publisher) Processor {
	return Processor{
		rabbitmqPublisher: rabbitmqPublisher,
	}
}

const (
	RountingKeyCreate = "create"
	RoutingKeyDelete  = "delete"
)

func (p *Processor) Process(d rabbitmq.Delivery) rabbitmq.Action {
	var err error
	switch string(d.RoutingKey) {
	case RountingKeyCreate:
		assignment := globalmodel.Assignment{}
		err = json.Unmarshal(d.Body, &assignment)
	case RoutingKeyDelete:
		logrus.Info("Delete")
	}

	if err != nil {
		err = p.rabbitmqPublisher.Publish(d.Body, []string{d.RoutingKey})
		if err != nil {
			logrus.Error("Publish failed: ", err)
		}
	}
	return rabbitmq.Ack
}
