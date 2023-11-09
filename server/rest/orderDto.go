package rest

type createOrderReq struct{
	OrderId string `json:"order_id"`
	ProductId string `json:"product_id"`
	CustomerId string `json:"customer_id"`
	Quantity int64 `json:"quantity"`
	TotalAmount float64 `json:"total_amount"`
}