package service

type Order struct {
	OrderID     string  `json:"OrderId"`
	CustomerID  string  `json:"CustomerId"`
	ProductID   string  `json:"ProductId"`
	Quantity    int     `json:"Quantity"`
	TotalAmount float64 `json:"TotalAmount"`
	OrderDate   int64   `json:"OrderDate"`
}

type OrdersResult struct {
	Orders []*Order `json:"Orders"`
}