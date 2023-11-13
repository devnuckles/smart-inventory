package service

type Order struct {
	ID           string  `json:"OrderId"`
	ProductName  string  `json:"ProductName"`
	BuyingPrice  float64 `json:"BuyingPrice"`
	Category     string  `json:"Category"`
	VendorEmail  string  `json:"VendorEmail"`
	Quantity     int     `json:"Quantity"`
	OrderDate    int64   `json:"OrderDate"`
	DeliveryDate int64   `json:"DeliveryDate"`
	Status       string  `json:"Status"`
}

type OrdersResult struct {
	Orders []*Order `json:"Orders"`
}
