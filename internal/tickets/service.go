package tickets

import (
	"context"
)

// The Service interface acts as a service layer for managing tickets.
type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

/*
The ServiceImpl struct is the implementation of the Service interface. It has methods
for manage tickets data.

	Fields:
	 - repository (Repository): An instance of the Repository interface.
*/
type ServiceImpl struct {
	repository Repository
}

/*
The NewService function returns a new instance of the ServiceImpl.

	Parameters:
	 - repository (Repository): An instance of the Repository interface.

	Returns:
	 - (Service): A new instance of the Service interface.
*/
func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

/*
The GetTotalTickets method returns the total number of available tickets in the "database" with a
certain destination country in it.

	Parameters:
	 - ctx (context.Context): The context of the request.
	 - destination (string): The country of destination.

	Returns:
	 - int: The total number of tickets with the specific destination.
	 - error: An error if there is no tickets available in the "database".
*/
func (s *ServiceImpl) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	ticketsByDestination, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(ticketsByDestination), nil
}

/*
The AverageDestination method returns the average number of tickets available in the "database" with
a certain destination country in it.

	Parameters:
	- ctx (context.Context): The context of the request.
	- destination (string): The country of destination.

	Returns:
	- float64: The average number of tickets available in the "database" with a certain destination country.
	- error: An error if there is no tickets available in the "database" or no tickets found with the given destination country.
*/
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
