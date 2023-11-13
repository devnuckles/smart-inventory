package rest

import (
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/service"
	"github.com/Tonmoy404/Smart-Inventory/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) createProduct(ctx *gin.Context) {
	var req CreateProductRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad request"))
		return
	}

	productID, err := uuid.NewUUID()
	if err != nil {
		logger.Error(ctx, "cannot generate productID", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
	}

	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		logger.Error(ctx, "cannot extract filename", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}
	defer file.Close()

	fileURL, err := s.svc.UploadFile(ctx, file, fileHeader)
	if err != nil {
		logger.Error(ctx, "cannot upload file into s3", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	createdAt := util.GetCurrentTimestamp()
	expiryDate := createdAt + +(req.ExpiryDate * 24 * 60 * 60 * 1000)

	product := &service.Product{
		ID:             productID.String(),
		Name:           req.Name,
		Quantity:       req.Quantity,
		Image:          fileURL,
		BuyingPrice:    req.BuyingPrice,
		Category:       req.Category,
		Status:         "available",
		ThreSholdValue: req.ThreSholdValue,
		CreatedAt:      createdAt,
		ExpiryDate:     expiryDate,
	}

	pro, err := s.svc.CreateProduct(ctx, product)
	if err != nil {
		logger.Error(ctx, "could not create product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}
	ctx.JSON(http.StatusCreated, s.svc.Response(ctx, "Product Created Successfully", pro))
}

func (s *Server) updateProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	var req UpdateProductReq
	err := ctx.ShouldBind(&req)
	if err != nil {
		logger.Error(ctx, "cannot pass validation", err)
		ctx.JSON(http.StatusBadRequest, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Bad Request"))
		return
	}

	product, err := s.svc.GetProductById(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	product.Name = req.Name
	product.BuyingPrice = req.BuyingPrice
	product.Quantity = req.Quantity
	product.ExpiryDate = req.ExpiryDate
	product.ThreSholdValue = req.ThreSholdValue
	product.Category = req.Category

	err = s.svc.UpdateProduct(ctx, product)
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

func (s *Server) getAllProducts(ctx *gin.Context) {
	productResult, err := s.svc.GetProducts(ctx)
	if err != nil {
		logger.Error(ctx, "cannot get all products", err)
		ctx.JSON(http.StatusInternalServerError, s.svc.Error(ctx, util.EN_INTERNAL_SERVER_ERROR, "Internal Server Error"))
		return
	}

	var products []*productResponse
	for _, pro := range productResult.Products {
		proRes := &productResponse{
			ID:          pro.ID,
			Name:        pro.Name,
			BuyingPrice: pro.BuyingPrice,
			Quantity:    pro.Quantity,
			ExpiryDate:  pro.ExpiryDate,
			CreatedAt:   pro.CreatedAt,
		}
		products = append(products, proRes)
	}

	allProducts := &getProductRes{
		Products: products,
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Successfully fetched all products", allProducts))
}

func (s *Server) getProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := s.svc.GetProductById(ctx, id)
	if err != nil {
		logger.Error(ctx, "cannot get product", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_NOT_FOUND, "Not Found"))
		return
	}

	if product == nil {
		logger.Error(ctx, "product not found", err)
		ctx.JSON(http.StatusNotFound, s.svc.Error(ctx, util.EN_API_PARAMETER_INVALID_ERROR, "Not Found"))
		return
	}

	getProduct := productRes{
		Name:        product.Name,
		Category:    product.Category,
		Image:       product.Image,
		BuyingPrice: product.BuyingPrice,
		Quantity:    product.Quantity,
		ExpiryDate:  product.ExpiryDate,
	}

	ctx.JSON(http.StatusOK, s.svc.Response(ctx, "Fetched product successfully", getProduct))
}
