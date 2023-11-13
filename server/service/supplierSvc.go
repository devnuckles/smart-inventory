package service

import (
	"context"
	"fmt"
)

func (s *service) GetSupplierByEmail(ctx context.Context, email string) (*Supplier, error) {
	user, err := s.supplierRepo.GetSupplierByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("cannot get user: %v", err)
	}

	return user, nil
}

func (s *service) GetSupplierByID(ctx context.Context, id string) (*Supplier, error) {
	user, err := s.supplierRepo.GetSupplierByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) CreateSupplier(ctx context.Context, sup *Supplier) error {
	err := s.supplierRepo.Create(ctx, sup)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteSupplier(ctx context.Context, id string) error {
	err := s.supplierRepo.DeleteSupplierByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateSupplier(ctx context.Context, user *Supplier) error {
	err := s.supplierRepo.UpdateSupplierByID(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetSuppliers(ctx context.Context) (*SupplierResult, error) {
	res, err := s.supplierRepo.GetAllSuppliers(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
