package service

import (
	"context"
	"fmt"
)

func (s *service) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.userRepo.GetItemByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("cannot get user: %v", err)
	}

	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, id string) (*User, error) {
	user, err := s.userRepo.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) CreateUser(ctx context.Context, user *User) error {
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteUser(ctx context.Context, id string) error {
	err := s.userRepo.DeleteItemByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateUser(ctx context.Context, user *User) error {
	err := s.userRepo.UpdateItemByID(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
