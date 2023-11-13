package rest

import (
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/Tonmoy404/Smart-Inventory/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) addSupplier(ctx *gin.Context) {
	var req addSupplier
	err := ctx.ShouldBind(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	supplier, err := s.svc.GetSupplierByEmail(ctx, req.Email)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	if supplier != nil {
		logger.Error(ctx, "already registered user with the email", req.Email)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_ALREADY_REGISTERED_ERROR, "Already registered with email"))
		return
	}

	ID, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot generate supplier id", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	supplier = &service.Supplier{
		ID:          ID.String(),
		Name:        req.Name,
		Product:     req.Product,
		Email:       req.Email,
		PhoneNumber: req.Contact,
	}

	err = s.svc.CreateSupplier(ctx, supplier)
	if err != nil {
		logger.Error(ctx, "cannot store user into db", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	ctx.JSON(http.StatusCreated, s.svc.Response(ctx, "Successfully created", nil))
}

func (s *Server) deleteSupplier(ctx *gin.Context) {
	id := ctx.Param("id")
	err := s.svc.DeleteSupplier(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot delete user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Deleted user successfully", nil))
}

func (s *Server) updateSuplier(ctx *gin.Context) {
	supId := ctx.Param("id")
	sup, err := s.svc.GetSupplierByID(ctx, supId)
	if err != nil {
		logger.Error(ctx, "cannot get supplier", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	if sup == nil {
		logger.Error(ctx, "supplier not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	var req updateSupplierReq
	err = ctx.ShouldBind(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal server error"))
		return
	}

	sup.Name = req.Name
	sup.Product = req.Product
	sup.Email = req.Email
	sup.PhoneNumber = req.Contact
	err = s.svc.UpdateSupplier(ctx, sup)
	if err != nil {
		logger.Error(ctx, "cannot update user", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Updated user successfully", nil))
}

func (s *Server) getAllSupplier(ctx *gin.Context) {
	suppliers, err := s.svc.GetSuppliers(ctx)
	if err != nil {
		logger.Error(ctx, "cannot get all suppliers", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	var supplierss []*supplierRes
	for _, sup := range suppliers.Suppliers {
		supRes := &supplierRes{
			ID:      sup.ID,
			Name:    sup.Name,
			Email:   sup.Email,
			Product: sup.Product,
			Contact: sup.PhoneNumber,
		}
		supplierss = append(supplierss, supRes)
	}

	allSuppliers := &getSupplyRes{
		Suppliers: supplierss,
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched all products", allSuppliers))

}

func (s *Server) getSupplier(ctx *gin.Context) {
	id := ctx.Param("id")
	supplier, err := s.svc.GetSupplierByID(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get Supplier", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	if supplier == nil {
		logger.Error(ctx, "user not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Not Found"))
		return
	}

	suppliers := supplierRes{
		ID: supplier.ID,
		Name: supplier.Name,
		Email: supplier.Email,
		Product: supplier.Product,
		Contact: supplier.PhoneNumber,

	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched user successfully", suppliers))
}
