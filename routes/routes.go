package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopher_rentals/controllers"
	"gopher_rentals/middleware"
)

func Routes() error {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/login", controllers.LoginController)
		api.POST("/register", controllers.RegisterController)

		customer := api.Group("/customers")
		{
			customer.GET("/", controllers.ListCustomersController).Use(middleware.Authorization())
			customer.GET("/:id", controllers.GetCustomerController).Use(middleware.Authorization())
			customer.POST("/create", controllers.CreateCustomerController).Use(middleware.Authorization())
			customer.PUT("/edit", controllers.EditCustomerController).Use(middleware.Authorization())
			customer.DELETE("/delete/:id", controllers.DeleteCustomerController).Use(middleware.Authorization())
		}

		car := api.Group("/cars")
		{
			car.GET("/", controllers.ListCarsController).Use(middleware.Authorization())
			car.GET("/:carId", controllers.GetCarController).Use(middleware.Authorization())
			car.POST("/create", controllers.CreateCarController).Use(middleware.Authorization())
			car.PUT("/edit", controllers.EditCarController).Use(middleware.Authorization())
			car.DELETE("/delete/:id", controllers.DeleteCarController).Use(middleware.Authorization())

			location := car.Group("/:carId/locations")
			{
				location.GET("/", controllers.ListCarLocationsController).Use(middleware.Authorization())
				location.GET("/:recent", controllers.ListCarRecentLocationsController).Use(middleware.Authorization())
				location.POST("/create", controllers.SaveCarLocationController).Use(middleware.Authorization())
				location.PUT("/edit", controllers.UpdateCarLocationController).Use(middleware.Authorization())
				location.DELETE("/delete/:id", controllers.DeleteCarLocationController).Use(middleware.Authorization())
			}
		}
	}

	err := router.Run(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server on port 8080 -> %v", err)
	}

	return nil
}
