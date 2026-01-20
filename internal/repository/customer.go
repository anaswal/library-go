package repository

import (
	"context"
	"library-go/domain"
)

type customerRepository struct{}

func NewCustomer() domain.CustomerRepository {
	return &customerRepository{}
}

func (cr customerRepository) FindAll(ctx context.Context) ([]domain.Customer, error) {
	panic("Implement")
}

func (cr customerRepository) FindById(ctx context.Context, id string) (domain.Customer, error) {
	panic("Implement")
}

func (cr customerRepository) Save(ctx context.Context, c *domain.Customer) error {
	panic("Implement")
}

func (cr customerRepository) Update(ctx context.Context, c *domain.Customer) {
	panic("Implement")
}

func (cr customerRepository) Delete(ctx context.Context, id string) error {
	panic("Implement")
}
