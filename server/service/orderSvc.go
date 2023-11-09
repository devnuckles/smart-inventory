package service

import "context"

func (s *service) CreateOrder(ctx context.Context, order *Order) error {
	err := s.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteOrder(ctx context.Context, id string) error {
	err := s.orderRepo.CancelOrderByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
