package service

import (
	globalmodel "github.com/GarnBarn/common-go/model"
	"github.com/GarnBarn/gb-assignment-consumer/repository"
)

type AssignmentService interface {
	CreateAssignment(assignment *globalmodel.Assignment) error
}

type assignmentService struct {
	assignmentRepository repository.AssignmentRepository
}

func (a *assignmentService) CreateAssignment(assignment *globalmodel.Assignment) error {
	return a.assignmentRepository.CreateAssignment(assignment)
}
