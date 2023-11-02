package rest

import (
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/Tonmoy404/Smart-Inventory/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) signupUser(ctx *gin.Context) {
	// validating req obj
	var req signupUserReq
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
		logger.Error(ctx, "already registered user with the email", nil)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_ALREADY_REGISTERED_ERROR, "Already registered"))
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
