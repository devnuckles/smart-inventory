package rest

import (
	"fmt"
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/config"
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

	router.GET("/api/test", server.test)
	router.POST("/api/users/signup", server.signupUser)
	router.POST("/api/users/login", server.loginUser)

	router.POST("/api/users/add", server.addUser)
	router.DELETE("/api/users/:id", server.deleteUser)
	router.GET("/api/users/:id", server.getUser)
	router.PATCH("/api/users/update", server.updateUser)
	router.POST("/api/users/logout", server.logoutUser)
	router.GET("/api/users/profile", server.getUserProfile)
	router.PATCH("/api/users/password", server.changePassword)

	///product routes

	router.POST("/api/item/create", server.createProduct)
	router.DELETE("/api/item/:id", server.deleteProduct)
	router.PATCH("/api/item/:id", server.updateProduct)
	router.GET("/api/items/all", server.getAllProducts)

	///order routes
	router.POST("/api/order/create", server.createOrder)
	router.DELETE("/api/order/:id", server.cancelOrder)
	router.GET("/api/order/:id", server.getOrder)
	router.PATCH("/api/order/:id", server.updateOrder)
	router.GET("/api/order/all", server.getAllOrders)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(fmt.Sprintf("%s:%s", server.appConfig.Host, server.appConfig.Port))
}

func (server *Server) test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "testing the server and its running successfully")
}
