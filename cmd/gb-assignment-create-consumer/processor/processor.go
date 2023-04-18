package processor

import (
	"encoding/json"

	globalmodel "github.com/GarnBarn/common-go/model"
	"github.com/GarnBarn/gb-assignment-consumer/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type Processor struct {
	rabbitmqPublisher    *rabbitmq.Publisher
	assignmentRepository repository.AssignmentRepository
}

func NewProcessor(rabbitmqPublisher *rabbitmq.Publisher, assignmentRepository repository.AssignmentRepository) Processor {
	return Processor{
		rabbitmqPublisher:    rabbitmqPublisher,
		assignmentRepository: assignmentRepository,
	}
}

const (
	RountingKeyCreate = "create"
	RoutingKeyDelete  = "delete"
)

func (p *Processor) Process(d rabbitmq.Delivery) error {
	assignment := globalmodel.Assignment{}
	err := json.Unmarshal(d.Body, &assignment)
	if err != nil {
		logrus.Error("Can't unmarshal data: ", err)
		return err
	}

	err = p.assignmentRepository.CreateAssignment(&assignment)
	if err != nil {
		logrus.Error("Can't save data: ", err)
		return err
	}

	logrus.Info("Successfully created the assignment id: ", assignment.ID)
	return nil
}
