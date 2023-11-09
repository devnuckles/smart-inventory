package rest

import (
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/Tonmoy404/Smart-Inventory/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) createOrder(ctx *gin.Context) {
	var req createOrderReq
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad Request"))
		return
	}

	orderId, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot create orderId", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	order := &service.Order{
		OrderID:     orderId.String(),
		ProductID:   req.ProductId,
		CustomerID:  req.CustomerId,
		Quantity:    int(req.Quantity),
		TotalAmount: req.TotalAmount,
	}

	err = s.svc.CreateOrder(ctx, order)
	if err != nil {
		logger.Error(ctx, "cannot create Order", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}


}
