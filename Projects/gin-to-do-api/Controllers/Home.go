package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
