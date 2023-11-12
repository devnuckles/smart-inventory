package service

type User struct {
	ID          string `json:"Id"`
	Username    string `json:"Username"`
	Fullname    string `json:"Fullname"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	PhoneNumber string `json:"PhoneNumber"`
	Image       string `json:"Picture"`
	Role        string `json:"Role"`
	Status      string `json:"Status"`
	CreatedAt   int64  `json:"CreatedAt"`
	CreatedBy   string `json:"CreatedBy"`
}
