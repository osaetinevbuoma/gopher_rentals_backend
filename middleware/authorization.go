package middleware

import (
	"github.com/gin-gonic/gin"
	"gopher_rentals/auth"
	"gopher_rentals/controllers"
	"log"
	"net/http"
	"strings"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		authorization := context.Request.Header.Get("Authorization")
		if authorization == "" {
			context.JSON(http.StatusForbidden, "No Authorization header")
			context.Abort()
			return
		}

		var token string
		t := strings.Split(authorization, "Bearer ")
		if len(t) == 2 {
			token = strings.TrimSpace(t[1])
		} else {
			context.JSON(http.StatusBadRequest, "Invalid authorization token format")
			context.Abort()
			return
		}

		maker, err := auth.NewJWTMaker(controllers.SecretKey)
		if err != nil {
			log.Println(err)
			return
		}

		payload, err := maker.VerifyToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, err.Error())
			context.Abort()
			return
		}

		context.Set("gopher_rentals_user", payload)
		context.Next()
	}
}
