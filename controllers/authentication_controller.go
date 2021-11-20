package controllers

import (
	"github.com/gin-gonic/gin"
	"gopher_rentals/auth"
	"gopher_rentals/services"
	"log"
)

type RegisterUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type LoginData struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func Register(context *gin.Context) {
	var registerUser RegisterUser

	err := context.ShouldBindJSON(&registerUser)
	if err != nil {
		log.Println(err)

		context.JSON(400, gin.H{"error": "Email, password and confirm are required"})
		context.Abort()
		return
	}

	if registerUser.Password != registerUser.ConfirmPassword {
		context.JSON(400, gin.H{"error": "Passwords do not match"})
		context.Abort()
		return
	}

	_, err = services.GetUser(registerUser.Email)
	if err != nil { // no user with the email address
		user, err := services.CreateUser(registerUser.Email, registerUser.Password)
		if err != nil {
			context.JSON(500, gin.H{"error": err})
		}

		context.JSON(200, user)
		return
	} else {
		context.JSON(409, gin.H{"error": "A user with this email already exists"})
		context.Abort()
		return
	}
}

func Login(context *gin.Context) {
	var login LoginData

	err := context.ShouldBindJSON(&login)
	if err != nil {
		context.JSON(400, gin.H{"error": "Email and password are required"})
		context.Abort()
		return
	}

	user, err := services.GetUser(login.Email)
	if err != nil {
		context.JSON(404, gin.H{"error": err})
		context.Abort()
		return
	}

	isOK := services.CheckPassword(user.Password, login.Password)
	if !isOK {
		context.JSON(404, gin.H{"error": "Wrong username/password combination"})
		context.Abort()
		return
	}

	secretKey := "adfasnasfasfg!#$!@#$!@FBTRURYREV#%^&$%_Y+RFSDV:<WRterteRGMKYOH#%$$#%$&#4543"
	maker, err := auth.NewJWTMaker(secretKey)
	if err != nil {
		log.Println(err)
		return
	}

	token, err := maker.CreateToken(login.Email, 60 * 60 * 24)
	if err != nil {
		log.Println(err)
		return
	}

	context.Header("Authorization", token)
	context.JSON(200, token)
	return
}
