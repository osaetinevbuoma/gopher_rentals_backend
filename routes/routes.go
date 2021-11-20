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
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
	}

	err := router.Run(":8080")
	if err != nil {
		return fmt.Errorf("failed to start server on port 8080 -> %v", err)
	}

	return nil
}
