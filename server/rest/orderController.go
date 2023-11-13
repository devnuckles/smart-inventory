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
		ID:           orderId.String(),
		ProductName:  req.ProductName,
		VendorEmail:  req.VendorEmail,
		Category:     req.Category,
		Quantity:     int(req.Quantity),
		OrderDate:    util.GetCurrentTimestamp(),
		DeliveryDate: req.DeliveryDate,
	}

	err = s.svc.CreateOrder(ctx, order)
	if err != nil {
		logger.Error(ctx, "cannot create Order", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	orderBody, _ := s.svc.GetOrderBody(ctx, order)
	mailBody := generateEmailBody()
	err = s.svc.SendMail(ctx, []string{req.VendorEmail}, "Request for New Order",mailBody , orderBody)

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Order Created Successfully", nil))
}

func (s *Server) cancelOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")

	err := s.svc.DeleteOrder(ctx, orderID)
	if err != nil {
		logger.Error(ctx, "could not delete order", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Order Deleted Successfully", nil))
}

func (s *Server) getOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	order, err := s.svc.GetOrderByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get order", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	if order == nil {
		logger.Error(ctx, "order not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Not Found"))
		return
	}

	getOrder := orderRes{
		ProductName:  order.ProductName,
		VendorName:   order.VendorEmail,
		Category:     order.Category,
		BuyingPrice:  order.BuyingPrice,
		Quantity:     int64(order.Quantity),
		OrderDate:    order.OrderDate,
		DeliveryDate: order.DeliveryDate,
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched order successfully", getOrder))
}

func (s *Server) updateOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")
	existingOrder, err := s.svc.GetOrderByID(ctx, orderID)
	if err != nil {
		logger.Error(ctx, "cannot get the order", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if existingOrder == nil {
		logger.Error(ctx, "order not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	var req updateOrderReq
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	order := &service.Order{
		ProductName:  req.ProductName,
		Category:     req.Category,
		BuyingPrice:  req.BuyingPrice,
		VendorEmail:  req.VendorEmail,
		Quantity:     int(req.Quantity),
		DeliveryDate: req.DeliveryDate,
	}

	err = s.svc.UpdateOrder(ctx, order)
	if err != nil {
		logger.Error(ctx, "could not update order", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Order Updated Successfully", order))
}

func (s *Server) getAllOrders(ctx *gin.Context) {
	orders, err := s.svc.GetAllOrders(ctx)
	if err != nil {
		logger.Error(ctx, "cannot get all orders", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Response(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	var orderRes []*orderResponse
	for _, order := range orders.Orders {
		res := &orderResponse{
			OrderId:      order.ID,
			ProductName:  order.ProductName,
			BuyingPrice:  order.BuyingPrice,
			Quantity:     int64(order.Quantity),
			DeliveryDate: order.DeliveryDate,
			Status:       order.Status,
		}
		orderRes = append(orderRes, res)
	}

	allOrders := &getAllOrderRes{
		Orders: orderRes,
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully get all orders", allOrders))
}
