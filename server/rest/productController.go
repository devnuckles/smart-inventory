package rest

import (
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/Tonmoy404/Smart-Inventory/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) CreateProduct(ctx *gin.Context) {
	var req CreateProductRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	productID, err := uuid.NewUUID()
	product := &service.Product{
		ID:          productID.String(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	pro, err := s.svc.CreateProduct(ctx, product)
	if err != nil {
		logger.Error(ctx, "could not create product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}
	ctx.JSON(http.StatusCreated, s.svc.Response(ctx, "Product Created Successfully", pro))
}

func (s *Server) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var req UpdateProductReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad Request"))
		return
	}

	updatedProduct := &service.Product{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}

	err = s.svc.UpdateProduct(ctx, updatedProduct)
	if err != nil {
		logger.Error(ctx, "cannot update product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Product Updated Successfully", nil))
}

func (s *Server) deleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.svc.DeleteProduct(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot delete product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Product user successfully", nil))
}
