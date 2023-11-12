package service

import (
	"context"
	"time"
)

type Service interface {
	//User Service
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, id string) error
	GetUserByID(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, user *User) error

	/////Product Services
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	UpdateProduct(ctx context.Context, updatedProduct *Product) error
	DeleteProduct(ctx context.Context, id string) error
	GetProducts(ctx context.Context) (*ProductsResult, error)

	////Order Services
	CreateOrder(ctx context.Context, order *Order) error
	DeleteOrder(ctx context.Context, id string) error
	GetOrderByID(ctx context.Context, id string) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) error
	GetAllOrders(ctx context.Context) (*OrdersResult, error)

	///Others
	Error(ctx context.Context, internalCode string, description string) *ErrorResponse
	Response(ctx context.Context, description string, data interface{}) *ResponseData
	IsPermitted(ctx context.Context, role, action, object string) bool
}

type UserRepo interface {
	Create(ctx context.Context, user *User) error
	GetItemByEmail(ctx context.Context, email string) (*User, error)
	DeleteItemByID(ctx context.Context, id string) error
	GetItemByID(ctx context.Context, id string) (*User, error)
	UpdateItemByID(ctx context.Context, user *User) error
}

type ProductRepo interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	UpdateProduct(ctx context.Context, product *Product) error
	DeleteProductById(ctx context.Context, id string) error
	GetAllProducts(ctx context.Context) (*ProductsResult, error)
}

type OrderRepo interface {
	CreateOrder(ctx context.Context, order *Order) error
	CancelOrderByID(ctx context.Context, id string) error
	GetOrder(ctx context.Context, id string) (*Order, error)
	UpdateOrder(ctx context.Context, order *Order) error
	GetAllItems(ctx context.Context) (*OrdersResult, error)
}

type ErrorRepo interface {
	GetError(ctx context.Context, internalCode string) (*ErrorDetail, error)
}

type FileRepo interface {
}

type Cache interface {
	Set(key string, value string, ttl time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	GetTTL(key string) (time.Duration, error)
}

type Authorization interface {
	IsPermitted(ctx context.Context, role, action, object string) bool
}
