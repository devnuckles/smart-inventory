package rest

type signupUserReq struct {
	Email           string `json:"email" binding:"required,email,max=100"`
	Password        string `json:"password" binding:"required,min=8,max=100"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type updateUserReq struct {
	Email       string `json:"email" binding:"required,email"`
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	PhoneNumber string `json:"phone_number"`
}

type updatePasswordReq struct {
	OldPassword     string `json:"old_password" binding:"required"`
	Password        string `json:"new_password" binding:"required,min=8,max=100"`
	ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
}

type addUserReq struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=8,max=100"`
	Role     string `json:"role" binding:"required,userRole"`
	Status   string `json:"status" binding:"required,userStatus"`
}

type loginUserReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type loginUserRes struct {
	Token string `json:"token"`
}

type userResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	PostCounts  int64  `json:"post_counts"`
	Status      string `json:"status"`
	CreatedAt   int64  `json:"created_at"`
}