package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (s *ServiceImpl) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	ticketsByDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(ticketsByDestination), nil
}

func (s *ServiceImpl) AverageDestination(ctx context.Context, destination string) (float64, error) {
	allTickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return 0, err
	}

	ticketsByDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	totalTickets := len(allTickets)
	totalTicketsByDestination := len(ticketsByDestination)
	return (float64(totalTicketsByDestination) / float64(totalTickets)) * 100, nil
}
