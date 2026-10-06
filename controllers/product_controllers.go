package controllers

import (
	"net/http"
	"strconv"
	"wms-prototype/services"
	"wms-prototype/views"

	"github.com/gin-gonic/gin"
)

func CreateProduct(C *gin.Context) {

	println("===== MASUK CONTROLLER CREATE PRODUCT =====")
	var productrequest views.ProductRequest
	var responsemodel views.BaseResponse

	err := C.ShouldBindJSON(&productrequest)
	if err != nil {
		responsemodel.Status = http.StatusBadRequest
		responsemodel.Error = append(responsemodel.Error, err.Error())
		C.JSON(http.StatusBadRequest, responsemodel)
		return
	}
	if err := services.CreateProduct(productrequest); err != nil {
		responsemodel.Status = http.StatusInternalServerError
		responsemodel.Error = append(responsemodel.Error, err.Error())

		C.JSON(http.StatusInternalServerError, responsemodel)
	}

	responsemodel.Status = http.StatusCreated
	responsemodel.Data = productrequest

	C.JSON(http.StatusCreated, responsemodel)
}

func GetProducts(c *gin.Context) {

	var response views.BaseResponse

	products, err := services.GetProducts()
	if err != nil {

		response.Status = http.StatusInternalServerError
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = products

	c.JSON(http.StatusOK, response)
}
func GetProductbyID(c *gin.Context) {

	var response views.BaseResponse

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		return
	}
	product, err := services.GetProductbyID(uint(id))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, "Product Not Found")

		c.JSON(http.StatusNotFound, response)
		return
	}
	response.Status = http.StatusOK
	response.Data = product

	c.JSON(http.StatusOK, response)
}

func UpdateProduct(c *gin.Context) {

	var response views.BaseResponse
	var request views.ProductRequest

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, "Product Not Found")

		c.JSON(http.StatusBadGateway, response)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = services.UpdateProduct(uint(id), request)
	if err != nil {
		response.Status = http.StatusNotFound
		response.Error = append(response.Error, "Product Not Found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = request

	c.JSON(http.StatusOK, response)

}

func DeleteProduct(c *gin.Context) {

	var response views.BaseResponse

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Status = http.StatusBadRequest
		response.Error = append(response.Error, "Invalid Product ID")
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = services.DeleteProduct(uint(id))
	if err != nil {
		response.Status = http.StatusNotFound
		response.Error = append(response.Error, "Product Not Found")
		c.JSON(http.StatusNotFound, response)
		return
	}

	response.Status = http.StatusOK
	response.Data = "Product Deleted Successfully"
	c.JSON(http.StatusOK, response)
}
