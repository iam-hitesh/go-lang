package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../Helpers"
	"../Models"
	"../Services"
)


func Home (req *gin.Context) {
	req.JSON(http.StatusOK, gin.H{
		"page": "Welcome to Home Page",
	})
}


func About (req *gin.Context) {
	req.JSON(http.StatusOK, gin.H{
		"page": "About page",
	})
}


func Login (req *gin.Context) {
	var user Models.User
	var existingUser Models.User

	if err := req.Bind(&user); err != nil {
		return
	}

	// Will check if any existing user or not and returns
	if err := existingUser.GetByEmail(user.Email); err != nil {
		req.JSON(http.StatusBadRequest, gin.H{
			"error": "User Doesn't exist",
		})
		return
	}

	if Helpers.VerifyPassword(existingUser.Password, user.Password) {
		req.JSON(http.StatusOK, gin.H{
			"token": Services.GenerateToken(&user),
		})
	} else {
		req.JSON(http.StatusOK, gin.H{
			"page": "Welcome to Home Page",
		})
	}
}
