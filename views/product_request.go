package views

type ProductRequest struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Location string `json:"location"`
}
