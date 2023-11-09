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

func (s *service) GetOrderByID(ctx context.Context, id string) (*Order, error) {
	order, err := s.orderRepo.GetOrder(ctx, id)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *service) UpdateOrder(ctx context.Context, order *Order) error {
	err := s.orderRepo.UpdateOrder(ctx, order)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetAllOrders(ctx context.Context) ([]*OrdersResult, error) {
	orders, err := s.orderRepo.GetAllItems(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
