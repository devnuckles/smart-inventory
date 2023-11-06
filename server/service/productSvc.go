package service

import "context"

func (s *service) CreateProduct(ctx context.Context, product *Product) (*Product, error) {
	product, err := s.productRepo.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, err
}
