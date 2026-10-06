package views

type BaseResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  []string    `json:"error"`
}
