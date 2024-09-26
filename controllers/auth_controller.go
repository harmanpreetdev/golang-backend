package controllers

import (
	"demo/database"
	"demo/models"
	"demo/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error creating user")
		return
	}

	user.Password = string(hashedPassword)

	if err := models.CreateUser(database.DB, &user); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Error creating user")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

var jwtSecretKey = []byte("secret_key")

func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request data")
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		utils.RespondWithError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}
