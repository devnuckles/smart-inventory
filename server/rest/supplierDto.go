package rest

type addSupplier struct {
	Name    string `form:"name" binding:"required"`
	Product string `form:"product" binding:"required"`
	Email   string `form:"email" binding:"required,email"`
	Contact string `form:"contact" binding:"required"`
}

type updateSupplierReq struct {
	Name    string `form:"name" binding:"required"`
	Product string `form:"product" binding:"required"`
	Email   string `form:"email" binding:"required,email"`
	Contact string `form:"contact" binding:"required"`
}

type supplierRes struct {
	ID      string `json:"id"`
	Name    string `form:"name" `
	Product string `form:"product" `
	Email   string `form:"email"`
	Contact string `form:"contact" `
}

type getSupplyRes struct {
	Suppliers []*supplierRes
}