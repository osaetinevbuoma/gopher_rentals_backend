package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gopher_rentals/services"
	"gopher_rentals/util"
	"net/http"
	"strconv"
)

func ListCarLocationsController(context *gin.Context) {
	locations, err := services.GetCarLocations(uuid.MustParse(context.Param("carId")))
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}
	
	context.JSON(http.StatusOK, locations)
}

func ListCarRecentLocationsController(context *gin.Context) {
	carId := uuid.MustParse(context.Param("carId"))
	recent := context.Param("recent")

	r, err := strconv.ParseInt(recent, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	locations, err := services.GetCarsRecentLocations(carId, int(r))
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, locations)
}

func SaveCarLocationController(context *gin.Context) {
	carId := uuid.MustParse(context.Param("carId"))
	data := util.UnmarshalJson(context)
	if data == nil {
		return
	}

	location, err := services.SaveCarLocation(carId, data)
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, location)
}

func UpdateCarLocationController(context *gin.Context) {
	carId := uuid.MustParse(context.Param("carId"))
	data := util.UnmarshalJson(context)
	if data == nil {
		return
	}

	location, err := services.UpdateCarLocation(carId, data)
	if err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, location)
}

func DeleteCarLocationController(context *gin.Context) {
	if err := services.DeleteCarLocation(uuid.MustParse(context.Param("id"))); err != nil {
		context.JSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, nil)
}
