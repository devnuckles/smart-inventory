package rest

type Payload struct {
	ID          string `json:"id"`
	UserName    string `json:"username"`
	FullName    string `json:"full_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	Role        string `json:"role"`
}
