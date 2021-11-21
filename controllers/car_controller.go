package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopher_rentals/models"
	"gopher_rentals/services"
	"net/http"
)

func ListCarsController(context *gin.Context) {
	cars, err := services.ListCars()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, cars)
}

func GetCarController(context *gin.Context) {
	car, err := services.GetCar(uuid.MustParse(context.Param("carId")))
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, car)
}

func CreateCarController(context *gin.Context) {
	var car models.Car

	if err := context.ShouldBindJSON(&car); err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	_, err := services.CreateCar(&car)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.JSON(http.StatusOK, car)
}

func EditCarController(context *gin.Context) {
	var car models.Car

	if err := context.ShouldBindJSON(&car); err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if _, err := services.UpdateCar(&car); err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, car)
}

func DeleteCarController(context *gin.Context) {
	if err := services.DeleteCar(uuid.MustParse(context.Param("id"))); err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, nil)
}
