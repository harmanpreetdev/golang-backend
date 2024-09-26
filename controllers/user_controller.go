package controllers

import (
	"demo/database"
	"demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUsers controller
func GetUsers(c *gin.Context) {
	users, err := models.GetAllUsers(database.DB) // Pass the DB connection here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser controller
func CreateUser(c *gin.Context) {	
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the database connection from the database package
	if err := models.CreateUser(database.DB, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}
	c.JSON(http.StatusCreated, user)
}
