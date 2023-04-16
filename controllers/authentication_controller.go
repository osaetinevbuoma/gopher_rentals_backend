package controllers

import (
	"github.com/gin-gonic/gin"
	"gopher_rentals/auth"
	"gopher_rentals/services"
	"log"
	"net/http"
)

var SecretKey string = "adfasnasfasfg!#$!@#$!@FBTRURYREV#%^&$%_Y+RFSDV:<WRterteRGMKYOH#%$$#%$&#4543"

type RegisterUser struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func RegisterController(context *gin.Context) {
	var registerUser RegisterUser

	err := context.ShouldBindJSON(&registerUser)
	if err != nil {
		log.Println(err)

		context.JSON(http.StatusBadRequest, gin.H{"error": "Email, password and confirm are required"})
		context.Abort()
		return
	}

	if registerUser.Password != registerUser.ConfirmPassword {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		context.Abort()
		return
	}

	_, err = services.GetUser(registerUser.Email)
	if err != nil { // no user with the email address
		user, err := services.CreateUser(registerUser.Email, registerUser.Password)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		context.JSON(http.StatusOK, user)
		return
	} else {
		context.JSON(http.StatusConflict, gin.H{"error": "A user with this email already exists"})
		context.Abort()
		return
	}
}

func LoginController(context *gin.Context) {
	var login LoginData

	err := context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		context.Abort()
		return
	}

	user, err := services.GetUser(login.Email)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err})
		context.Abort()
		return
	}

	isOK := services.CheckPassword(user.Password, login.Password)
	if !isOK {
		context.JSON(http.StatusNotFound, gin.H{"error": "Wrong username/password combination"})
		context.Abort()
		return
	}

	maker, err := auth.NewJWTMaker(SecretKey)
	if err != nil {
		log.Println(err)
		return
	}

	token, err := maker.CreateToken(user.ID, login.Email, 24*7)
	if err != nil {
		log.Println(err)
		return
	}

	context.Header("Authorization", token)
	context.JSON(http.StatusOK, token)
	return
}
