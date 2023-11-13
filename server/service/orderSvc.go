package service

import (
	"context"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

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

func (s *service) GetAllOrders(ctx context.Context) (*OrdersResult, error) {
	orders, err := s.orderRepo.GetAllItems(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (s *service) GetOrderBody(ctx context.Context, order *Order) (*excelize.File, error) {
	file := excelize.NewFile()

	headers := []string{"Order ID", "Product Name", "Vendor Email", "Category", "Quantity", "Order Date", "Delivery Date"}
	values := []string{order.ID, order.ProductName, order.VendorEmail, order.Category, fmt.Sprint(order.Quantity), string(order.OrderDate), string(order.DeliveryDate)}

	for col, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+col, 1)
		file.SetCellValue("Sheet1", cell, header)
	}

	for col, value := range values {
		cell := fmt.Sprintf("%c%d", 'A'+col, 2)
		file.SetCellValue("Sheet1", cell, value)
	}

	return file, nil
}
