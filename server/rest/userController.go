package rest

import (
	"net/http"
	"time"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/Tonmoy404/Smart-Inventory/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) signupUser(ctx *gin.Context) {
	// validating req obj
	var req signupUserReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	// check existing user by email
	// user, err := s.svc.GetUserByEmail(ctx, req.Email)
	// if err != nil {
	// 	logger.Error(ctx, "cannot get user", err)
	// 	ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
	// 	return
	// }

	// if user != nil {
	// 	logger.Error(ctx, "already registered user with the email", nil)
	// 	ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_ALREADY_REGISTERED_ERROR, "Already registered"))
	// 	return
	// }

	// hash the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.salt.SecretKey)
	if err != nil {
		logger.Error(ctx, "cannot hash the password", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	// create random user id
	userID, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot generate user id", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	// create the user into db
	user := &service.User{
		ID:          userID.String(),
		Username:    "Not Set",
		Fullname:    "Not Set",
		Email:       req.Email,
		Password:    string(hashedPass),
		Role:        "user",
		PhoneNumber: "Not Set",
		Status:      "approved",
		CreatedAt:   util.GetCurrentTimestamp(),
	}

	err = s.svc.CreateUser(ctx, user)
	if err != nil {
		logger.Error(ctx, "cannot store user into db", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	ctx.JSON(http.StatusCreated, s.svc.Response(ctx, "Successfully created", nil))
}

func (s *Server) addUser(ctx *gin.Context) {
	// validating req obj
	var req addUserReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	// check existing user by email
	user, err := s.svc.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Error(ctx, "cannot get user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	if user != nil {
		logger.Error(ctx, "already registered user with the email", req.Email)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_ALREADY_REGISTERED_ERROR, "Already registered with email"))
		return
	}

	// hash the password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.salt.SecretKey)
	if err != nil {
		logger.Error(ctx, "cannot hash the password", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	// create random user id
	userID, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot generate user id", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	// create the user into db
	user = &service.User{
		ID:          userID.String(),
		Username:    "Not Set",
		Fullname:    "Not Set",
		Email:       req.Email,
		Password:    string(hashedPass),
		Role:        req.Role,
		PhoneNumber: "Not Set",
		Status:      req.Status,
		CreatedAt:   util.GetCurrentTimestamp(),
	}

	err = s.svc.CreateUser(ctx, user)
	if err != nil {
		logger.Error(ctx, "cannot store user into db", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	ctx.JSON(http.StatusCreated, s.svc.Response(ctx, "Successfully created", nil))
}

func (s *Server) loginUser(ctx *gin.Context) {
	var req loginUserReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	user, err := s.svc.GetUserByEmail(ctx, req.Email)
	if err != nil {
		logger.Error(ctx, "cannot get user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	if user == nil {
		logger.Error(ctx, "user not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	// if user.Status != "approved" {
	// 	logger.Error(ctx, "Forbidden", err)
	// 	ctx.JSON(http.StatusForbidden, s.svc.Response(ctx, util.EN_UNAUTHORIZED_ERROR, "Forbidden"))
	// 	return
	// }

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), ([]byte(req.Password)))
	if err != nil {
		logger.Error(ctx, "cannot decrypt the password", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Invalid Credentials"))
		return
	}

	// create access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":           user.ID,
		"username":     user.Username,
		"full_name":    user.Fullname,
		"email":        user.Email,
		"phone_number": user.PhoneNumber,
		"role":         user.Role,
		"status":       user.Status,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	accessToken, err := token.SignedString([]byte(s.jwt.SecretKey))
	if err != nil {
		logger.Error(ctx, "failed to generate token", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	res := loginUserRes{
		Token: accessToken,
	}

	ctx.SetCookie("token", accessToken, 3600, "/", "", false, true)
	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "successfully logged in", res))
}

func (s *Server) deleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.svc.DeleteUser(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot delete user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Deleted user successfully", nil))
}

func (s *Server) updateUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(Payload)
	user, err := s.svc.GetUserByID(ctx, authPayload.ID)
	if err != nil {
		logger.Error(ctx, "cannot get user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if user == nil {
		logger.Error(ctx, "user not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	var req updateUserReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	user.Username = req.Username
	user.Fullname = req.Fullname
	user.Email = req.Email
	user.PhoneNumber = req.PhoneNumber

	err = s.svc.UpdateUser(ctx, user)
	if err != nil {
		logger.Error(ctx, "cannot update user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Updated user successfully", nil))
}

func (s *Server) logoutUser(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.Status(http.StatusOK)
}

func (s *Server) getUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user, err := s.svc.GetUserByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get user", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	if user == nil {
		logger.Error(ctx, "user not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Not Found"))
		return
	}

	userRes := userResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched user successfully", userRes))
}

func (s *Server) getUserProfile(ctx *gin.Context) {
	// decode cookie
	authPayload := ctx.MustGet(authorizationPayloadKey).(Payload)

	user, err := s.svc.GetUserByID(ctx, authPayload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	loggedInUserData := userResponse{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Fullname:    user.Fullname,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		Status:      user.Status,
		CreatedAt:   user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Logged in user data", loggedInUserData))
}

func (s *Server) changePassword(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(Payload)

	user, err := s.svc.GetUserByID(ctx, authPayload.ID)
	if err != nil {
		logger.Error(ctx, "cannot get user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if user == nil {
		logger.Error(ctx, "user not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	var req updatePasswordReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), ([]byte(req.OldPassword)))
	if err != nil {
		logger.Error(ctx, "cannot decrypt the password", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Invalid Password"))
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), s.salt.SecretKey)
	if err != nil {
		logger.Error(ctx, "cannot hash the password", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	user.Password = string(hashedPass)

	err = s.svc.UpdateUser(ctx, user)
	if err != nil {
		logger.Error(ctx, "cannot update user password", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Updated Password Successfully", nil))
}
