package rest

import (
	"fmt"
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/config"
	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router    *gin.Engine
	appConfig *config.Application
	svc       service.Service
	jwt       *config.Token
	salt      *config.Salt
}

func NewServer(appConfig *config.Application, svc service.Service, salt *config.Salt, jwt *config.Token) (*Server, error) {
	server := &Server{
		appConfig: appConfig,
		svc:       svc,
		salt:      salt,
		jwt:       jwt,
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// CORS middleware
	router.Use(corsMiddleware)

	// log middleware
	router.Use(logger.ModifyContext)

	router.GET("/api/test", server.test)
	router.POST("/api/users/signup", server.signupUser)
	router.POST("/api/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(server.authMiddleware())

	authRoutes.POST("/api/users/add", server.addUser)
	authRoutes.DELETE("/api/users/:id", server.deleteUser)
	authRoutes.GET("/api/users/:id", server.getUser)
	authRoutes.PATCH("/api/users/update", server.updateUser)
	authRoutes.POST("/api/users/logout", server.logoutUser)
	authRoutes.GET("/api/users/profile", server.getUserProfile)
	authRoutes.PATCH("/api/users/password", server.changePassword)

	///product routes

	authRoutes.POST("/api/item/create", server.createProduct)
	authRoutes.DELETE("/api/item/:id", server.deleteProduct)
	authRoutes.PATCH("/api/item/:id", server.updateProduct)
	authRoutes.GET("/api/items/all", server.getAllProducts)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(fmt.Sprintf("%s:%s", server.appConfig.Host, server.appConfig.Port))
}

func (server *Server) test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "testing the server and its running successfully")
}
