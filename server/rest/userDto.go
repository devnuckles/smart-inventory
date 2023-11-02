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