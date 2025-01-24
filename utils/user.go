package utils

import (
	"ynbauth/models"
	
	"net/http"
	"github.com/gin-gonic/gin"
)

// Route for creating a new account
func  RegisterUser(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		FullName string `json:"full_name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Call the CreateUser function to save the user
	err := models.CreateUser(user.Username, user.FullName, user.Email,user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User account successfully created."})
}


