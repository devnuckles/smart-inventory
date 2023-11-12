package rest

type CreateProductRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int64   `json:"quantity" binding:"required"`
	ExpiryDate  int64   `json:"expiry_date"`
}

type CreateProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

type UpdateProductReq struct {
	Name        string  `json:"name" binding:"required"`
	BuyingPrice float64 `json:"buying_price" binding:"required"`
	Quantity    int64   `json:"quantity" binding:"required"`
	ExpiryDate  int64   `json:"expiry_date" binding:"required"`
}

type productResponse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	BuyingPrice float64 `json:"buying_price"`
	Quantity    int64   `json:"quantity"`
	ExpiryDate  int64   `json:"expiry_date"`
	CreatedAt   int64   `json:"created_at"`
}

type getProductRes struct {
	Products []*productResponse
}
