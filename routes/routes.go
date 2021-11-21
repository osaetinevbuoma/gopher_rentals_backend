package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopher_rentals/controllers"
)

func Routes() error {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/login", controllers.LoginController)
		api.POST("/register", controllers.RegisterController)

		customer := api.Group("/customers")
		{
			customer.GET("/", controllers.ListCustomersController)
			customer.GET("/:id", controllers.GetCustomerController)
			customer.POST("/create", controllers.CreateCustomerController)
			customer.PUT("/edit", controllers.EditCustomerController)
			customer.DELETE("/delete/:id", controllers.DeleteCustomerController)
		}

		car := api.Group("/cars")
		{
			car.GET("/", controllers.ListCarsController)
			car.GET("/:carId", controllers.GetCarController)
			car.POST("/create", controllers.CreateCarController)
			car.PUT("/edit", controllers.EditCarController)
			car.DELETE("/delete/:id", controllers.DeleteCarController)

			location := car.Group("/:carId/locations")
			{
				location.GET("/", controllers.ListCarLocationsController)
				location.GET("/:recent", controllers.ListCarRecentLocationsController)
				location.POST("/create", controllers.SaveCarLocationController)
				location.PUT("/edit", controllers.UpdateCarLocationController)
				location.DELETE("/delete/:id", controllers.DeleteCarLocationController)
			}
		}
	}

	err := router.Run(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server on port 8080 -> %v", err)
	}

	return nil
}
