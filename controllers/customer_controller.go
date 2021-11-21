package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopher_rentals/models"
	"gopher_rentals/services"
	"gopher_rentals/util"
	"net/http"
)

func ListCustomersController(context *gin.Context) {
	customers, err := services.ListCustomers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, customers)
}

func GetCustomerController(context *gin.Context) {
	customer, err := services.GetCustomer(uuid.MustParse(context.Param("id")))
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, customer)
}

func CreateCustomerController(context *gin.Context) {
	jsonData := util.UnmarshalJson(context)
	if jsonData == nil {
		return
	}

	customer, err := services.CreateCustomer(jsonData)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, customer)
}

func EditCustomerController(context *gin.Context) {
	var customer models.Customer

	err := context.ShouldBindJSON(&customer)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	_, err = services.EditCustomer(&customer)
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, customer)
}

func DeleteCustomerController(context *gin.Context) {
	err := services.DeleteCustomer(uuid.MustParse(context.Param("id")))
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, nil)
}