package Controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"../Models"
)


func CreateUser(req *gin.Context) {
	var user Models.User

	req.Bind(&user)

	err := Models.CreateUser(&user)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, user)
	}
}


func GetAllUsers(req *gin.Context) {
	var users []Models.User

	err := Models.GetAllUsers(&users)

	if err != nil {
		req.AbortWithStatus(http.StatusBadRequest)
	} else {
		req.JSON(http.StatusOK, users)
	}
}
