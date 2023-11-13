package rest

type createOrderReq struct {
	ProductName  string `json:"product_name" binding:"required"`
	BuyingPrice  string `json:"buying_price" binding:"required"`
	VendorEmail  string `json:"vendor_email" binding:"required"`
	Category     string `json:"category" binding:"required"`
	Quantity     int64  `json:"quantity" binding:"required"`
	DeliveryDate int64  `json:"delivery_date" binding:"required"`
}

type updateOrderReq struct {
	ProductName  string  `json:"product_name" binding:"required"`
	VendorEmail  string  `json:"vendor_email" binding:"required"`
	Category     string  `json:"category" binding:"category"`
	Quantity     int64   `json:"quantity" binding:"required"`
	BuyingPrice  float64 `json:"total_amount" binding:"required"`
	DeliveryDate int64   `json:"delivery_date" binding:"required"`
}

type orderRes struct {
	ProductName  string  `json:"product_name" `
	VendorName   string  `json:"vendor_name" `
	BuyingPrice  float64 `json:"buying_price"`
	Category     string  `json:"category"`
	Quantity     int64   `json:"quantity" `
	OrderDate    int64   `json:"order_date"`
	DeliveryDate int64   `json:"delivery_date"`
}

type orderResponse struct {
	OrderId      string  `json:"order_id"`
	ProductName  string  `json:"product_name"`
	BuyingPrice  float64 `json:"buying_price"`
	Quantity     int64   `json:"quantity"`
	DeliveryDate int64   `json:"delivery_date"`
	Status       string  `json:"status"`
}

type getAllOrderRes struct {
	Orders []*orderResponse `json:"orders"`
}
