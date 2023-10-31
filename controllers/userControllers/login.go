package userControllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"taskung.com/test/m/v2/inits"
	"taskung.com/test/m/v2/models"
)

func Login(c *gin.Context) {
	var user models.User
	var existUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check exist Username ?
	inits.DB.Where("username = ?", user.Username).First(&existUser)
	if existUser.ID == 0 {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "Username does not exists"})
		return
	}

	// Find user by username
	inits.DB.Where("username = ?", user.Username).First(&existUser)

	// Compare the provided password with the hashed password
	err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Invalid credentials, Login failed",
			"message": "Wrong password",
		})
		return
	}

	// Generate JWT token
	token := generateToken(user.ID)

	c.JSON(http.StatusOK, gin.H{
		"user_id":  existUser.ID,
		"status":   "ok",
		"message":  "login successful",
		"username": user.Username,
		"token":    token,
	})

}
