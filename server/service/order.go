package service

type Order struct {
	OrderID     string  `json:"OrderId"`
	CustomerID  string  `json:"CustomerID"`
	ProductID   string  `json:"ProductID"`
	Quantity    int     `json:"Quantity"`
	TotalAmount float64 `json:"TotalAmount"`
	OrderDate   int64   `json:"OrderDate"`
}
