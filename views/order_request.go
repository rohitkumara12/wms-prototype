package views

type OrderItemRequest struct {
	ProductID uint `json:"productid" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}

type OrderRequest struct {
	OrderNumber string             `json:"ordernumber" binding:"required"`
	Items       []OrderItemRequest `json:"items" binding:"required"`
}
