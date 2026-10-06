package views

type StockMovementRequest struct {
	ProductID uint `json:"productid" binding:"required"`
	Quantity  uint `json:"quantity" binding:"required"`
}
