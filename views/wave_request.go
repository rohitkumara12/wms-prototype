package views

type WaveRequest struct {
	WaveNumber string `json:"wavenumber" binding:"required"`
	OrderIDs   []uint `json:"orderids" binding:"required,min=1"`
}

type WaveStatusRequest struct {
	Status string `json:"status" binding:"required"`
}
type WaveResponse struct {
	WaveNumber string   `json:"wavenumber"`
	Status     string   `json:"status"`
	Orders     []string `json:"orders"`
}
