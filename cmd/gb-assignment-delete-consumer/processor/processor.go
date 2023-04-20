package processor

import (
	"encoding/json"

	globalmodel "github.com/GarnBarn/common-go/model"
	rabbitMQ "github.com/GarnBarn/common-go/rabbitmq"
	"github.com/GarnBarn/gb-assignment-consumer/pkg/repository"
	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type Processor struct {
	rabbitmqPublisher    *rabbitmq.Publisher
	assignmentRepository repository.AssignmentRepository
}

func NewProcessor(rabbitmqPublisher *rabbitmq.Publisher, assignmentRepository repository.AssignmentRepository) rabbitMQ.Processor {
	return Processor{
		rabbitmqPublisher:    rabbitmqPublisher,
		assignmentRepository: assignmentRepository,
	}
}

const (
	RountingKeyCreate = "create"
	RoutingKeyDelete  = "delete"
)

func (p Processor) Process(d rabbitmq.Delivery) error {
	assignment := globalmodel.AssignmentDeleteRequest{}
	err := json.Unmarshal(d.Body, &assignment)
	if err != nil {
		logrus.Error("Can't unmarshal data: ", err)
		return err
	}

	err = p.assignmentRepository.DeleteAssignment(assignment.ID)
	if err != nil {
		logrus.Error("Can't delete data: ", err)
		return err
	}

	logrus.Info("Successfully deleted the assignment id: ", assignment.ID)
	return nil
}
