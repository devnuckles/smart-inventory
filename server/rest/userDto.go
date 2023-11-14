package rest

type signupUserReq struct {
	Name     string `json:"name" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// type updateUserReq struct {
// 	Email       string `json:"email" binding:"required,email"`
// 	Username    string `json:"username"`
// 	Fullname    string `json:"fullname"`
// 	PhoneNumber string `json:"phone_number"`
// }

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

type updateUserReq struct {
	Email       string `form:"email" binding:"required,email"`
	Username    string `form:"username" binding:"required"`
	Fullname    string `form:"fullname" binding:"required"`
	PhoneNumber string `form:"phone_number" binding:"required"`
}
