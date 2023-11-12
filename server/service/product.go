package service

type Product struct {
	ID             string  `json:"Id"`
	Name           string  `json:"Name"`
	BuyingPrice    float64 `json:"BuyingPrice"`
	Quantity       int64   `json:"Quantity"`
	ThreSholdValue int64   `json:"ThresholdValue"`
	ExpiryDate     int64   `json:"ExpiryDate"`
	Status         string  `json:"Status"`
	CreatedAt      int64   `json:"CreatedAt"`
}

type ProductsResult struct {
	Products []*Product `json:"Products"`
}
