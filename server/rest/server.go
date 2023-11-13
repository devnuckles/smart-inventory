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

	authRoutes := router.Group("/").Use(server.authMiddleware())

	authRoutes.POST("/api/users/add", server.addUser)
	authRoutes.DELETE("/api/users/:id", server.deleteUser)
	authRoutes.GET("/api/users/:id", server.getUser)
	authRoutes.PATCH("/api/users/update", server.updateUser)
	authRoutes.POST("/api/users/logout", server.logoutUser)
	authRoutes.GET("/api/users/profile", server.getUserProfile)
	authRoutes.PATCH("/api/users/password", server.changePassword)

	///product routes

<<<<<<< HEAD
	authRoutes.POST("/api/items/create", server.createProduct) //testing done
	authRoutes.DELETE("/api/items/:id", server.deleteProduct)  //testing done
	authRoutes.PATCH("/api/items/:id", server.updateProduct)   //testing done
	authRoutes.GET("/api/items/:id", server.getProduct)        //testing done
	authRoutes.GET("/api/items/all", server.getAllProducts)    //testing done
=======
	authRoutes.POST("/api/item/screate", server.createProduct)
	authRoutes.DELETE("/api/item/s:id", server.deleteProduct)
	authRoutes.PATCH("/api/items/:id", server.updateProduct)
	authRoutes.GET("/api/items/all", server.getAllProducts)
>>>>>>> 04a3369 (Rewfactor)

	///order routes
	authRoutes.POST("/api/orders/create", server.createOrder)
	authRoutes.DELETE("/api/orders/:id", server.cancelOrder)
	authRoutes.GET("/api/orders/:id", server.getOrder)
	authRoutes.PATCH("/api/orders/:id", server.updateOrder)
	authRoutes.GET("/api/orders/all", server.getAllOrders)

	///supplier routes

	authRoutes.POST("/api/suppliers/add", server.addSupplier)
	authRoutes.DELETE("/api/suppliers/:id", server.deleteSupplier)
	authRoutes.PATCH("/api/suppliers/:id", server.updateSuplier)
	authRoutes.GET("/api/suppliers/all", server.getAllSupplier)
	authRoutes.GET("/api/suppliers/:id", server.getSupplier)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(fmt.Sprintf("%s:%s", server.appConfig.Host, server.appConfig.Port))
}

func (server *Server) test(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "testing the server and its running successfully")
}
