package service

type Supplier struct {
	ID          string `json:"Id"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
	Product     string `json:"Product"`
	Return      string `json:"Return"`
	Quantity    int64  `json:"Quantity"`
}
