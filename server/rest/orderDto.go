package rest

type createOrderReq struct {
	ProductId   string  `json:"product_id" binding:"required"`
	CustomerId  string  `json:"customer_id" binding:"required"`
	Quantity    int64   `json:"quantity" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
}

type updateOrderReq struct {
	ProductId   string  `json:"product_id" binding:"required"`
	CustomerId  string  `json:"customer_id" binding:"required"`
	Quantity    int64   `json:"quantity" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required"`
}

type orderRes struct {
	ProductId   string  `json:"product_id" `
	CustomerId  string  `json:"customer_id" `
	Quantity    int64   `json:"quantity" `
	TotalAmount float64 `json:"total_amount" `
}

type orderResponse struct {
	OrderId     string  `json:"order_id"`
	ProductID   string  `json:"product_id"`
	CustomerID  string  `json:"customer_id"`
	Quantity    int64   `json:"quantity"`
	TotalAmount float64 `json:"total_amount"`
}

type getAllOrderRes struct{
	Orders []*orderResponse `json:"orders"`
}