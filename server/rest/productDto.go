package rest

type CreateProductRequest struct {
	Name           string  `form:"name" binding:"required"`
	BuyingPrice    float64 `form:"buying_price" binding:"required"`
	Category       string  `form:"category" binding:"required"`
	Quantity       int64   `form:"quantity" binding:"required"`
	ExpiryDate     int64   `form:"expiry_date"`
	ThreSholdValue int64   `form:"threshold_value" binding:"required"`
}

type CreateProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

type UpdateProductReq struct {
	Name           string  `form:"name" binding:"required"`
	Category       string  `form:"category" binding:"required"`
	BuyingPrice    float64 `form:"buying_price" binding:"required"`
	Quantity       int64   `form:"quantity" binding:"required"`
	ThreSholdValue int64   `form:"threshold_value" binding:"required"`
	ExpiryDate     int64   `form:"expiry_date" binding:"required"`
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

type productRes struct {
	Name        string  `json:"name"`
	Category    string  `json:"category"`
	Image       string  `json:"img"`
	BuyingPrice float64 `json:"buying_price"`
	Quantity    int64   `json:"quantity"`
	ExpiryDate  int64   `json:"expiry_date"`
}
