package service

type Product struct {
	ID          string  `json:"Id"`
	Name        string  `json:"Name"`
	Description string  `json:"Description"`
	Price       float64 `json:"Price"`
	Quantity    int64   `json:"Quantity"`
	CreatedAt   int64   `json:"CreatedAt"`
}

type ProductsResult struct {
	Products []*Product `json:"Products"`
}
